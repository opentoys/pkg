package webx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXxx(t *testing.T) {
	root := newTrie()

	root.GET("/", func(ctx context.Context) error {
		return nil
	})

	root.Use(func(ctx context.Context) (e error) {
		fmt.Println("all use")
		return
	})

	root.GET("/api/foo/", func(ctx context.Context) {
		fmt.Println("/api/foo")
	})

	root.GET("/api/foo", func(ctx context.Context) {
		fmt.Println("/api/foo")
	})

	api := root.Group("/api/")
	api.Use(func(ctx context.Context) (e error) {
		fmt.Println("api use")
		return
	})

	api.GET("/foo", func(ctx context.Context) error {
		return nil
	})

	api.GET("/#(?<name>[a-z]{1,}).html", func(ctx context.Context) error {
		fmt.Println(
			GetRequest(ctx).PathValue(""),
			GetRequest(ctx).PathValue("name"),
		)
		return nil
	})

	api.POST("/*", func(ctx context.Context) error {
		fmt.Println("* ignore")
		return nil
	})

	api.POST("/:abc", func(ctx context.Context) error {
		fmt.Println("* abc")
		return nil
	})

	api.POST("/*/bar", func(ctx context.Context) error {
		fmt.Println("* bar")
		return nil
	})

	api.POST("/abc", func(ctx context.Context) error {
		fmt.Println("POST /api/abc")
		return nil
	})

	fmt.Println(root)
	// http.Handle("/", root)

	testhttp(root, "GET", "/?abc=123", bytes.NewBufferString(`[{"Name":"sajhawghd"}]`))
	testhttp(root, "POST", "/api/asdasdasd?abc=123", bytes.NewBufferString(`[{"Name":"sajhawghd"}]`))
	testhttp(root, "POST", "/api/abc?abc=123", bytes.NewBufferString(`[{"Name":"sajhawghd"}]`))
	testhttp(root, "POST", "/api/asdasdasd/bar?abc=123", bytes.NewBufferString(`[{"Name":"sajhawghd"}]`))
}

func testhttp(mux http.Handler, method, uri string, body io.Reader) (w *httptest.ResponseRecorder) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	w = httptest.NewRecorder()
	r := httptest.NewRequest(method, uri, body)
	r.Header.Add("Content-Type", "application/json;charset=utf8")
	r = r.WithContext(ctx)
	mux.ServeHTTP(w, r)
	return
}
