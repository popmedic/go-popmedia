package server

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type search struct {
	searchIndex      map[string]string
	indexMutex       sync.RWMutex
	createMutex      sync.Mutex
	setCreatingMutex sync.RWMutex
	creating         bool
}

var (
	searchInstance *search
	searchOnce     sync.Once
)

func MainSearch() *search {
	searchOnce.Do(func() {
		searchInstance = newSearch()
	})
	return searchInstance
}

func newSearch() *search {
	s := &search{}
	s.setCreating(false)
	go s.indexRoutine()

	return s
}

func (s *search) setSearchIndex(idx map[string]string) {
	func(s *search, v map[string]string, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.Unlock()
		}(lock)
		lock.Lock()
		s.searchIndex = v
	}(s, idx, &s.indexMutex)
}

func (s *search) getSearchIndex() map[string]string {
	v := make(map[string]string)
	func(s *search, v *map[string]string, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.RUnlock()
		}(lock)
		lock.RLock()
		*v = s.searchIndex
	}(s, &v, &s.indexMutex)
	return v
}

func (s *search) getSearchValues() []string {
	idx := s.getSearchIndex()
	ss := []string{}
	if !s.getCreating() {
		for _, v := range idx {
			if len(v) > 0 && v[0] != '.' && v[0] != '_' {
				ss = append(ss, v)
			}
		}
	}
	return ss
}

func (s *search) setCreating(v bool) {
	func(s *search, v bool, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.Unlock()
		}(lock)
		lock.Lock()
		s.creating = v
	}(s, v, &s.setCreatingMutex)
}

func (s *search) getCreating() bool {
	var v bool
	func(s *search, v *bool, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.RUnlock()
		}(lock)
		lock.RLock()
		*v = s.creating
	}(s, &v, &s.setCreatingMutex)
	return v
}

func (s *search) createIndex() {
	if !s.getCreating() {
		defer s.setCreating(false)
		s.setCreating(true)
		log.Println("creating search index...")
		idx := map[string]string{}
		i := 0
		fmt.Println()
		out := ""
		err := Walk(MainConfig.Root, func(path string, info os.FileInfo, err error) error {
			if nil != info {
				if !strings.HasPrefix(".", info.Name()) && !strings.HasPrefix("_", info.Name()) {
					if info.IsDir() ||
						stringsContain(MainConfig.MediaExt, strings.ToLower(filepath.Ext(info.Name()))) {

						idx[path] = strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
						i++
						for range out {
							fmt.Print("\b")
						}
						out = fmt.Sprintf("processed %d files", i)
						fmt.Print(out)
					}
				}
			}
			return nil
		})
		fmt.Println()
		if nil != err {
			log.Println("unable to create index", err)
		} else {
			log.Println("search index created.")
			s.setSearchIndex(idx)
		}
	}
}

func (s *search) Query(v string) map[string]string {
	idx := s.getSearchIndex()
	res := map[string]string{}
	re, err := regexp.Compile(strings.ToLower(".*" + v + ".*"))
	if nil != err {
		log.Println(err)
		return res
	}
	for key, val := range idx {
		if re.MatchString(strings.ToLower(val)) {
			res[key] = val
		}
	}
	return res
}

func (s *search) indexRoutine() {
	s.createIndex()
	for {
		select {
		case <-time.After(time.Hour * 4):
			s.createIndex()
		}
	}
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

// walk recursively descends path, calling w.
func walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	info, p := followSymLinkWithPath(path, info)
	err := walkFn(path, info, nil)
	if err != nil {
		if info.IsDir() && err == filepath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() ||
		strings.HasPrefix(".", info.Name()) ||
		strings.HasPrefix("_", info.Name()) {
		return nil
	}
	names, err := readDirNames(p)
	if err != nil {
		return walkFn(path, info, err)
	}

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil && err != filepath.SkipDir {
				return err
			}
		} else {
			err = walk(filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != filepath.SkipDir {
					return err
				}
			}
		}
	}
	return nil
}

func Walk(root string, walkFn filepath.WalkFunc) error {
	info, err := os.Lstat(root)
	if err != nil {
		return walkFn(root, nil, err)
	}
	return walk(root, info, walkFn)
}

func init() {
	MainSearch()
}
