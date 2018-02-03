package dir

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

//drwxr-xr-x  2 Kevin  staff  68 Feb  3 08:59 dir1ls -n
var (
	pemsRegx = regexp.MustCompile("drwxr-xr-x  2 Kevin  staff  68 Feb  3 08:59 dir1")
)

func LS(p string) ([]string, error) {
	out, err := exec.Command("ls", p).Output()
	if nil != err {
		return nil, err
	}
	res := strings.Split(strings.Trim(string(out), "\n"), "\n")
	if len(res) == 1 && res[0] == "" {
		return []string{}, nil
	}
	return res, nil
}

func IsDir(p string) bool {
	info, err := os.Stat(p)
	if nil == err {
		return info.IsDir()
	}
	return false
}

func LSFiles(p string) []string {
	res := []string{}
	Walk(p, func(pth string) {
		res = append(res, pth)
	})
	return res
}

func Walk(p string, do func(pth string)) {
	l, e := LS(p)
	if nil == e && len(l) > 0 {
		for _, pp := range l {
			fp := filepath.Join(p, pp)
			if IsDir(fp) {
				Walk(fp, do)
			} else {
				do(fp)
			}
		}
	}
}

type bbufCloser struct {
	*bytes.Buffer
}

func (b *bbufCloser) Close() error {
	return nil
}
