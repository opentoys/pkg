package webx_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/opentoys/webx"
)

type Login struct {
	Name string `json:"n" qs:"name" header:"X-Name"`
	Abc  []int  `form:"abc"`
}

type ErrCode struct {
	Msg  string
	Code int
}

func (e *ErrCode) Error() string {
	return e.Msg
}

var app webx.Route

func CustomSend(c context.Context, a any, e error) webx.H {
	if e != nil {
		var w = webx.GetResponse(c)
		switch vv := e.(type) {
		case *ErrCode:
			return webx.H{"code": vv.Code, "msg": vv.Msg}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return webx.H{"code": -1, "msg": e.Error()}
		}
	}

	return webx.H{"code": 0, "msg": "ok", "data": a}
}

func Recover(ctx context.Context, a any) error {
	fmt.Println(a)
	return nil
}

func init() {
	app = webx.New(webx.WithCustomeSend(CustomSend), webx.WithRecover(Recover))
}

type Customer struct {
	Name string
}

func (s *Customer) Hello(ctx context.Context, in *Login) (hello string, e error) {
	if in.Name != "" {
		hello = in.Name + s.Name
		return
	}
	e = &ErrCode{Code: 123, Msg: "asjhadj"}
	return
}

func (s *Customer) Call(ctx context.Context, name []Login) (e error) {
	fmt.Println("call name", name)
	return
}

type ContextStore struct {
	session string
}

var wsonce = make(map[string]sync.Once)

func TestHTTP2(t *testing.T) {
	var wsmap sync.Map

	app.GET("/hello", func(ctx context.Context) (e error) {
		return errors.New("hello error")
	})

	app.GET("/ws/:platform/:id", func(ctx context.Context) (e error) {
		w := webx.GetOriginResponse(ctx)
		r := webx.GetRequest(ctx)
		var platform = r.PathValue("platform")
		var id = r.PathValue("id")
		defer webx.Abort(ctx)
		var key = "id" + platform + id
		_, ok := wsmap.Load(key)
		if ok {
			return
		}
		w.Write([]byte("hhh"))
		return errors.New("hello error")
	})

	srv := http.Server{Handler: app, Addr: ":8123"}
	t.Fatal(srv.ListenAndServe())
}

func TestWeb(t *testing.T) {
	a := big.NewFloat(0.1).SetPrec(512)
	b := big.NewFloat(0.2).SetPrec(512)
	fmt.Println(big.NewFloat(0).SetPrec(512).Add(b, a).Float64())

	var s = &Customer{}
	fmt.Println(app)

	app.Use(
		webx.UseLogger(func(s string, a ...any) {
			fmt.Printf(s, a...)
		}),
		webx.UsePayload(),
		func(ctx context.Context) (e error) {
			webx.Abort(ctx)
			return
		},
	)

	app.Use(func(ctx context.Context) (e error) {
		w := httptest.NewRecorder()
		webx.SetResponse(ctx, w)
		buf := webx.GetRawBody(ctx)
		fmt.Println("request start", string(buf))
		go func() {
			<-ctx.Done()
			fmt.Println("request end", w.Body.String(), w.Code)
		}()
		return nil
	})

	api := app.Group("/v1/api")

	api.Use(func(ctx context.Context) (e error) {
		// var r = webx.GetRequest(ctx)
		// var key = r.URL.Query().Get("key")
		// 根据 key 拉取加密配置
		// 解析 body
		var data = make(map[string]any)
		// 解密 body

		// 重新赋值
		var out = webx.GetPayload(ctx)
		if out != nil && data["body"] != nil {
			e = json.Unmarshal([]byte(data["body"].(string)), &out)
		}
		return
	})

	customer := api.Group("/customer", func(ctx context.Context) (e error) {
		r := webx.GetRequest(ctx)
		if r.URL.Query().Get("hello") == "" {
			return errors.New("hello must be")
		}
		return
	})
	customer.POST("/hello", s.Hello)
	customer.POST("/call", s.Call)

	merchant := api.Group("/merchant")
	merchant.POST("/login", s.Call)

	w := testhttp(app, "POST", "/v1/api/merchant/login?abc=123", bytes.NewBufferString(`[{"Name":"sajhawghd"}]`))

	fmt.Println(w.Body.String())
	// app.Router.ServeHTTP()
	// http.ListenAndServe("", app)

	<-time.After(time.Second)
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
