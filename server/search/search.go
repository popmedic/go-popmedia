package search

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/search/dir"
)

type Search struct {
	searchIndex      map[string]string
	indexMutex       sync.RWMutex
	createMutex      sync.Mutex
	setCreatingMutex sync.RWMutex
	creating         bool
}

var (
	searchInstance *Search
	searchOnce     sync.Once
)

func MainSearch() *Search {
	searchOnce.Do(func() {
		searchInstance = newSearch()
	})
	return searchInstance
}

func newSearch() *Search {
	s := &Search{}
	s.setCreating(false)
	go s.indexRoutine()

	return s
}

func (s *Search) setSearchIndex(idx map[string]string) {
	func(s *Search, v map[string]string, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.Unlock()
		}(lock)
		lock.Lock()
		s.searchIndex = v
	}(s, idx, &s.indexMutex)
}

func (s *Search) getSearchIndex() map[string]string {
	v := make(map[string]string)
	func(s *Search, v *map[string]string, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.RUnlock()
		}(lock)
		lock.RLock()
		*v = s.searchIndex
	}(s, &v, &s.indexMutex)
	return v
}

func (s *Search) getSearchValues() []string {
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

func (s *Search) setCreating(v bool) {
	func(s *Search, v bool, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.Unlock()
		}(lock)
		lock.Lock()
		s.creating = v
	}(s, v, &s.setCreatingMutex)
}

func (s *Search) getCreating() bool {
	var v bool
	func(s *Search, v *bool, lock *sync.RWMutex) {
		defer func(lock *sync.RWMutex) {
			lock.RUnlock()
		}(lock)
		lock.RLock()
		*v = s.creating
	}(s, &v, &s.setCreatingMutex)
	return v
}

func (s *Search) createIndex() {
	if !s.getCreating() {
		s.setCreating(true)
		defer s.setCreating(false)
		log.Println("creating search index...")
		idx := map[string]string{}
		i := 0
		out := ""
		dir.Walk(config.MainConfig.Root, func(p string) {
			b := filepath.Base(p)
			if !strings.HasPrefix("_", b) && stringsContain(config.MainConfig.MediaExt, filepath.Ext(b)) {
				idx[p] = strings.TrimSuffix(b, filepath.Ext(b))
				i++
				for range out {
					fmt.Print("\b")
				}
				out = fmt.Sprintf("processed %d files", i)
				fmt.Print(out)
			}
		})
		for range out {
			fmt.Print("\b")
		}
		log.Println("Total processed", i)
		log.Println("search index created.")
		s.setSearchIndex(idx)
	}
}

func (s *Search) Query(v string) map[string]string {
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

func (s *Search) indexRoutine() {
	s.createIndex()
	for {
		select {
		case <-time.After(time.Hour * 4):
			s.createIndex()
		}
	}
}

func stringsContain(ss []string, s string) bool {
	for _, str := range ss {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}
