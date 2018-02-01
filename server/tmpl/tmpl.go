package tmpl

import (
	"html/template"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
)

func LoadTemplate(name, path string) (*template.Template, error) {
	b, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}
	funcMap := template.FuncMap{
		"replaceDashes": replaceDashes,
		"splitDashes":   splitDashes,
		"split":         strings.Split,
		"joinPath":      joinPath,
		"stripPlayer":   stripPlayer,
		"upDir":         upDir,
		"urlEncode":     urlEncode,
	}
	return template.New(name).Funcs(funcMap).Parse(string(b))
}

func replaceDashes(s string) template.HTML {
	return template.HTML(strings.Replace(s, " - ", "<br/>", -1))
}

func splitDashes(s string) []string {
	return strings.Split(s, " - ")
}

func joinPath(ss []string, n int) string {
	return string(filepath.Separator) + filepath.Join(ss[0:n+1]...)
}

func stripPlayer(s string) string {
	return strings.Replace(s, "/player", "", 1)
}

func upDir(s string) string {
	p := filepath.SplitList(filepath.Dir(s))
	if len(p) > 1 {
		pp := filepath.Join(p[:len(p)-1]...)
		if pp[0] != '/' {
			return "/" + pp
		}
		return pp
	}
	return "/"
}

func urlEncode(s string) string {
	if len(s) > 0 && s[0] == '/' {
		return "/" + url.PathEscape(s[1:])
	}
	return url.PathEscape(s)
}
