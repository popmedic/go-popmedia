package rokuFeed

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/popmedic/go-popmedia/server/roku_feed/model"
)

var (
	//ErrorMsgNotADirectory - message returned by Generate when dirname is not a directory
	ErrorMsgNotADirectory = "not a directory"
)

// RokuFeed - Used for walking directories and creating the RokuFeed.json
type RokuFeed struct {
	dirname       string
	root          *model.Root
	mutexMovie    sync.RWMutex
	mutexPlaylist sync.RWMutex
	mutexSeries   sync.RWMutex
}

// NewRokuFeed - Creates a RokuFeed object off and sets the directory
func NewRokuFeed(dirname string) *RokuFeed {
	return &RokuFeed{
		dirname: dirname,
		root: &model.Root{
			Movies:          model.Movies{},
			Series:          model.Series{},
			ShortFormVideos: model.ShortFormVideos{},
			TvSpecials:      model.TvSpecials{},
			Categories:      model.Categories{},
			Playlists:       model.Playlists{},
		},
		mutexMovie:    sync.RWMutex{},
		mutexPlaylist: sync.RWMutex{},
		mutexSeries:   sync.RWMutex{},
	}
}

// Generate - Generates a RokuFeed as a JSON string
//  returns - an error on failure or a JSON string as a Roku feed of the directory on success
func (feed *RokuFeed) Generate() (string, error) {
	if info, err := os.Stat(feed.dirname); nil != err {
		return "", err
	} else if !info.IsDir() {
		return "", errors.New(ErrorMsgNotADirectory)
	} else {
		feed.root = &model.Root{
			ProviderName: "Popmedia",
			LastUpdated:  nowRFC3339(),
			Language:     "en",
			Movies:       model.Movies{},
			Series:       model.Series{},
			Playlists:    model.Playlists{},
		}
		var path string
		if path, err = filepath.Abs(info.Name()); nil != err {
			return "", err
		}
		if err := feed.walk(info, path); nil != err {
			return "", err
		}
		return "", nil
	}
}

func (feed *RokuFeed) walk(info os.FileInfo, path string) (err error) {
	if info.IsDir() &&
		!filepath.HasPrefix(filepath.Base(info.Name()), ".") &&
		!filepath.HasPrefix(filepath.Base(info.Name()), "_") {
		fmt.Println(">" + path)

		var infos []os.FileInfo
		if infos, err = ioutil.ReadDir(path); nil != err {
			return
		}
		for _, nfo := range infos {
			if err = feed.walk(nfo, filepath.Join(path, nfo.Name())); nil != err {
				return
			}
		}
		return
	}
	if isMp4(info) {
		fmt.Println("*" + filepath.Base(removeExt(info.Name())))
	}
	return
}

// MARK: Synchronised Functions

func (feed *RokuFeed) syncAddMovie(movie model.Movie) {
	feed.mutexMovie.Lock()
	defer feed.mutexMovie.Unlock()

	feed.root.Movies = append(feed.root.Movies, movie)
}

func (feed *RokuFeed) syncAddPlaylist(playlist model.Playlist) {
	feed.mutexPlaylist.Lock()
	defer feed.mutexPlaylist.Unlock()

	feed.root.Playlists = append(feed.root.Playlists, playlist)
}

func (feed *RokuFeed) syncContainsPlaylist(name string) *model.Playlist {
	feed.mutexPlaylist.RLock()
	defer feed.mutexPlaylist.RUnlock()

	for _, playlist := range feed.root.Playlists {
		if playlist.Name == name {
			return &playlist
		}
	}
	return nil
}

func (feed *RokuFeed) syncAddSeries(serie model.Serie) {
	feed.mutexSeries.Lock()
	defer feed.mutexSeries.Unlock()

	feed.root.Series = append(feed.root.Series, serie)
}

func (feed *RokuFeed) syncContainsSeries(title string) *model.Serie {
	feed.mutexSeries.RLock()
	defer feed.mutexSeries.RUnlock()

	for _, serie := range feed.root.Series {
		if serie.Title == title {
			return &serie
		}
	}
	return nil
}

// MARK: Helpers

func nowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func isMp4(fi os.FileInfo) bool {
	return strings.ToLower(filepath.Ext(fi.Name())) == ".mp4"
}

func removeExt(path string) string {
	return strings.TrimRight(path, filepath.Ext(path))
}
