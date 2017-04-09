package server

import "testing"

func TestNewConfigEmptyPath(t *testing.T) {
	_, err := newConfig("")
	if err == nil {
		t.Error("empty path should have failed")
	}
}

func TestNewConfigBadJSON(t *testing.T) {
	_, err := newConfig("config_bad.json")
	if err == nil {
		t.Error("bad json should have failed")
	}
}

func TestNewConfigSuccess(t *testing.T) {
	_, err := newConfig("config.json")
	if err != nil {
		t.Error("good config should have succeeded")
	}
}