package tmpl

import (
	"testing"
)

func TestUpDirName(t *testing.T) {
	type row struct {
		given string
		exp   string
	}
	tbl := []row{
		{
			given: "up/one/dir/",
			exp:   "one",
		},
		{
			given: "up/one/dir/test.go",
			exp:   "dir",
		},
		{
			given: "up/one/dir",
			exp:   "one",
		},
		{
			given: "up/one/",
			exp:   "up",
		},
		{
			given: "/up/",
			exp:   "Home",
		},
		{
			given: "/up",
			exp:   "Home",
		},
		{
			given: "/",
			exp:   "Home",
		},
	}
	for _, r := range tbl {
		if res := upDirName(r.given); res != r.exp {
			t.Errorf("updir test: given %q, expected %q, got %q", r.given, r.exp, res)
		}
	}
}
func TestUpDir(t *testing.T) {
	type row struct {
		given string
		exp   string
	}
	tbl := []row{
		{
			given: "up/one/dir/",
			exp:   "/up/one",
		},
		{
			given: "up/one/dir/test.go",
			exp:   "/up/one/dir",
		},
		{
			given: "up/one/dir",
			exp:   "/up/one",
		},
		{
			given: "up/one/",
			exp:   "/up",
		},
		{
			given: "/up/",
			exp:   "/",
		},
		{
			given: "/up",
			exp:   "/",
		},
		{
			given: "/",
			exp:   "/",
		},
	}
	for _, r := range tbl {
		if res := upDir(r.given); res != r.exp {
			t.Errorf("updir test: given %q, expected %q, got %q", r.given, r.exp, res)
		}
	}
}
