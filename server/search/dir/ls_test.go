package dir

// func TestLS(t *testing.T) {
// 	type row struct {
// 		given string
// 		exp   []string
// 	}
// 	tt := []row{
// 		{
// 			given: "notapath",
// 			exp:   nil,
// 		},
// 		{
// 			given: "testdir",
// 			exp:   []string{"dir1", "dir2", "file1", "testlink"},
// 		},
// 		{
// 			given: "testdir/dir1",
// 			exp:   []string{},
// 		},
// 	}
// 	for _, r := range tt {
// 		if res, err := LS(r.given); nil != err {
// 			if r.exp != nil {
// 				t.Error(err)
// 			}
// 		} else if !stringArray(res).Compare(r.exp) {
// 			t.Errorf("given %q expected %q actual %q",
// 				r.given,
// 				stringArray(r.exp).String(),
// 				stringArray(res).String())
// 		}
// 	}
// }

// func TestIsDir(t *testing.T) {
// 	type row struct {
// 		given string
// 		exp   bool
// 	}
// 	tt := []row{
// 		{
// 			given: "notapath",
// 			exp:   false,
// 		},
// 		{
// 			given: "testdir",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/dir1",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/dir2",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/file1",
// 			exp:   false,
// 		},
// 		{
// 			given: "testdir/testlink",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/testlink/linkdir1",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/testlink/linkdir2",
// 			exp:   true,
// 		},
// 		{
// 			given: "testdir/testlink/linkfile1",
// 			exp:   false,
// 		},
// 	}
// 	for _, r := range tt {
// 		if res := IsDir(r.given); res != r.exp {
// 			t.Errorf("given %q expected %t actual %t", r.given, r.exp, res)
// 		}
// 	}
// }

// func TestLSFiles(t *testing.T) {
// 	type row struct {
// 		given string
// 		exp   []string
// 	}
// 	tt := []row{
// 		{
// 			given: "testdir",
// 			exp:   []string{"testdir/file1", "testdir/testlink/linkfile1"},
// 		},
// 		{
// 			given: "testlink",
// 			exp:   []string{"testlink/linkfile1"},
// 		},
// 	}
// 	for _, r := range tt {
// 		if res := LSFiles(r.given); !stringArray(r.exp).Compare(res) {
// 			t.Errorf("given %q expected %q actual %q",
// 				r.given,
// 				stringArray(r.exp).String(),
// 				stringArray(res).String())
// 		}
// 	}
// }

// type stringArray []string

// func (sa stringArray) Contains(s string) bool {
// 	for _, str := range sa {
// 		if str == s {
// 			return true
// 		}
// 	}
// 	return false
// }
// func (sa1 stringArray) Compare(sa2 stringArray) bool {
// 	for _, s1 := range sa1 {
// 		if !sa2.Contains(s1) {
// 			return false
// 		}
// 	}
// 	for _, s2 := range sa2 {
// 		if !sa1.Contains(s2) {
// 			return false
// 		}
// 	}
// 	return true
// }
// func (sa stringArray) String() string {
// 	return strings.Join(sa, ", ")
// }
