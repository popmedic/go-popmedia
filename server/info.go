package server

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Info struct {
	Name  string
	Path  string
	Image string
	Desc  string
}

type InfoList []*Info

func (il InfoList) Len() int {
	return len(il)
}

func (il InfoList) Less(i, j int) bool {
	return strings.TrimPrefix(strings.ToLower(il[i].Name), "the ") < strings.TrimPrefix(strings.ToLower(il[j].Name), "the ")
}

// Swap swaps the elements with indexes i and j.
func (il InfoList) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func newInfo(name, path string) *Info {
	info := Info{
		Name: name,
		Path: path,
	}
	info.Image = strings.TrimSuffix(path, filepath.Ext(path)) + "-SD.jpg"
	if _, err := os.Stat(filepath.Join(MainConfig.Root, info.Image)); err != nil {
		info.Image = MainConfig.DirectoryImage
	}

	p := filepath.Join(MainConfig.Root, strings.TrimSuffix(path, filepath.Ext(path))+".desc")
	if _, err := os.Stat(p); err == nil {
		d, err := ioutil.ReadFile(p)
		if nil == err {
			info.Desc = string(d)
		}
	}
	return &info
}

type FilesAndDirectoriesInfo struct {
	Info        *Info
	Files       InfoList
	Directories InfoList
}

func (v *FilesAndDirectoriesInfo) SortAll() {
	sort.Sort(v.Files)
	sort.Sort(v.Directories)
}
