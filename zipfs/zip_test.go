package zipfs_test

import (
	"easyadmin/pkg/github/zipfs"
	"fmt"
	"net/http"
	"testing"
)

func TestXxx(t *testing.T) {
	http.DefaultServeMux.Handle("/", http.FileServerFS(zipfs.New("a.zip")))
	http.DefaultServeMux.Handle("/", zipfs.New("a.zip"))
}

func TestZip(t *testing.T) {
	f := zipfs.New("../../../pkg.zip")
	fmt.Println(f.List("json"))
	fmt.Println(f.Open("pkg/variable.go"))
	fmt.Println(f.ReadDir("pkg/helper/"))
}
