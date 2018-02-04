package handle

import (
	"net/http"
	"net/url"
	"testing"
)

func TestIs(t *testing.T) {
	type row struct {
		given string
		exp   bool
	}
	tt := []row{
		{
			given: "/images/foo.jpg",
			exp:   true,
		},
		{
			given: "/image/foo.jpg",
			exp:   false,
		},
		{
			given: "/images/foo.jp",
			exp:   false,
		},
		{
			given: "/images/foo.pnG",
			exp:   true,
		},
	}
	handle := NewImages()
	for i, r := range tt {
		req := &http.Request{
			URL: &url.URL{
				Path: r.given,
			},
		}
		if res := handle.Is(req); res != r.exp {
			t.Errorf("[row %d] given %q expected %t got %t", i+1, r.given, r.exp, res)
		}
	}
}
