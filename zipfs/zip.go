package zipfs

import (
	"archive/zip"
	"bytes"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Option func(*FS)

type FS struct {
	r      *zip.Reader
	err    error
	filter func(string) bool
}

func New(file string, ops ...Option) (fs *FS) {
	fs = &FS{}
	buf, e := os.ReadFile(file)
	if e != nil {
		fs.err = e
		return
	}
	return newBuffer(fs, buf, ops...)
}

func NewBuffer(buf []byte, ops ...Option) (fs *FS) {
	return newBuffer(fs, buf, ops...)
}

func NewReader(r io.Reader, ops ...Option) (fs *FS) {
	fs = &FS{}
	buf, e := io.ReadAll(r)
	if e != nil {
		fs.err = e
		return
	}
	return newBuffer(fs, buf, ops...)
}

func newBuffer(fs *FS, buf []byte, ops ...Option) *FS {
	if fs == nil {
		fs = &FS{}
	}
	for _, v := range ops {
		v(fs)
	}
	fs.r, fs.err = zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	return fs
}

func (s *FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.filter != nil && s.filter(r.URL.Path) {
		return
	}
	http.FileServerFS(s).ServeHTTP(w, r)
	return
}

func (s *FS) Open(name string) (fs.File, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.r.Open(name)
}

func (s *FS) List(v any) (lst []fs.DirEntry, e error) {
	lst = make([]fs.DirEntry, 0, len(s.r.File))
	var search, _ = v.(string)
	var reg, _ = v.(*regexp.Regexp)
	for _, v := range s.r.File {
		if search != "" {
			if strings.Contains(v.Name, search) {
				lst = append(lst, &zipdirifle{v: v})
			}
			continue
		}

		if reg != nil {
			if reg.Match([]byte(v.Name)) {
				lst = append(lst, &zipdirifle{v: v})
			}
			continue
		}

		lst = append(lst, &zipdirifle{v: v})
	}
	return
}

func (s *FS) ReadDir(dir string) (lst []fs.DirEntry, e error) {
	var lastidx = len(dir) - 1
	if dir[lastidx] != filepath.Separator {
		dir += "/"
	}
	for _, v := range s.r.File {
		if s.dirfile(v.Name, dir) {
			lst = append(lst, &zipdirifle{v: v})
		}
	}
	return
}

func (s *FS) dirfile(name, dir string) (ok bool) {
	var lastidx = len(name) - 1
	if name[lastidx] == filepath.Separator {
		name = name[:lastidx]
	}
	after, ok := strings.CutPrefix(name, dir)
	if !ok || after == "" {
		return false
	}
	ok = len(strings.Split(after, string(filepath.Separator))) == 1
	return
}

func (s *FS) ReadFile(name string) (buf []byte, e error) {
	if s.err != nil {
		return nil, s.err
	}
	f, e := s.r.Open(name)
	if e != nil {
		return
	}

	return io.ReadAll(f)
}

type zipdirifle struct {
	v *zip.File
}

func (s *zipdirifle) Name() string {
	return s.v.Name
}

func (s *zipdirifle) IsDir() bool {
	return s.v.FileInfo().IsDir()
}

func (s *zipdirifle) Type() fs.FileMode {
	return s.v.Mode()
}

func (s *zipdirifle) Info() (fs.FileInfo, error) {
	return s.v.FileInfo(), nil
}
