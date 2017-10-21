package rokuFeed

import (
	"os"
	"testing"
	"time"

	"github.com/popmedic/popmedia2/server/roku_feed/model"
)

func TestNewRokuFeed(t *testing.T) {
	if nil == NewRokuFeed("") {
		t.Error("Should have returned new RokuFeed object.")
	}
}

func TestGenerateErrorDirnameDoesNotExist(t *testing.T) {
	feed := NewRokuFeed("jfghkljfhklj")
	if _, err := feed.Generate(); nil == err {
		t.Error("should have returned error")
	} else if !os.IsNotExist(err) {
		t.Error("should have been a NotExist error:", err)
	}
}

func TestGenerateErrorDirnameNotADirectory(t *testing.T) {
	feed := NewRokuFeed(os.Args[0])
	if _, err := feed.Generate(); nil == err {
		t.Error("should have returned error")
	} else if err.Error() != ErrorMsgNotADirectory {
		t.Error("should have been \""+ErrorMsgNotADirectory+"\" error:", err)
	}
}

func TestGenerateSuccess(t *testing.T) {
	feed := NewRokuFeed("test_data") //filepath.Join(filepath.Dir(os.Args[0]), "test_data"))
	if res, err := feed.Generate(); nil != err {
		t.Error("should not have returned error:", err)
	} else if res != "" {
		t.Error("should have been \"\"")
	}
}

func TestSyncAddMovieSuccess(t *testing.T) {
	feed := NewRokuFeed("test_data")
	movieTitlesToAdd := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, movieTitleToAdd := range movieTitlesToAdd {
		go feed.syncAddMovie(model.Movie{Title: movieTitleToAdd})
	}
	time.Sleep(time.Millisecond * 2)
	for _, movie := range feed.root.Movies {
		found := false
		for _, movieTitleToAdd := range movieTitlesToAdd {
			if movie.Title == movieTitleToAdd {
				found = true
				break
			}
		}
		if !found {
			t.Error(movie.Title + " should have been added but was not found.")
		}
	}
}

func TestSyncAddSeriesSuccess(t *testing.T) {
	feed := NewRokuFeed("test_data")
	serieTitlesToAdd := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, serieTitleToAdd := range serieTitlesToAdd {
		go feed.syncAddSeries(model.Serie{Title: serieTitleToAdd})
	}
	time.Sleep(time.Millisecond * 2)
	for _, title := range serieTitlesToAdd {
		if s := feed.syncContainsSeries(title); nil == s {
			t.Error(title + " should have been added but was not found.")
		}
	}
}

func TestSyncAddPlaylistSuccess(t *testing.T) {
	feed := NewRokuFeed("test_data")
	playlistTitlesToAdd := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, playlistTitleToAdd := range playlistTitlesToAdd {
		go feed.syncAddPlaylist(model.Playlist{Name: playlistTitleToAdd})
	}
	time.Sleep(time.Millisecond * 2)
	for _, name := range playlistTitlesToAdd {
		if p := feed.syncContainsPlaylist(name); nil == p {
			t.Error(name + " should have been added but was not found.")
		}
	}
}
