package webx

import (
	"context"
	"net/http"
)

const StatusUnavailable = 499 // 自定义状态码

type H = map[string]any

type Handle func(context.Context) (e error)

type Route interface {
	http.Handler
	Use(fn ...Handle)
	Group(uri string, fn ...Handle) Route
	Any(method, uri string, handle ...any)
	GET(uri string, handle ...any)
	POST(uri string, handle ...any)
	PUT(uri string, handle ...any)
	PATCH(uri string, handle ...any)
	DELETE(uri string, handle ...any)
	OPTIONS(uri string, handle ...any)
	HEAD(uri string, handle ...any)
}

type Option func(any)

func WithCustomeSend(fn func(context.Context, any, error) H) Option {
	return func(s any) {
		switch v := s.(type) {
		case *Trie:
			v.custom = fn
		case *Router:
			v.customsend = fn
		}
	}
}

func WithDefaultContentType(typ string) Option {
	return func(s any) {
		switch v := s.(type) {
		case *Trie:
			v.defaultContentType = typ
		case *Router:
			v.defaultContentType = typ
		}
	}
}

func WithRecover(fn func(context.Context, any) error) Option {
	return func(s any) {
		switch v := s.(type) {
		case *Trie:
			v.recover = fn
		case *Router:
			v.recover = fn
		}
	}
}

func New(opts ...Option) Route {
	s := newTrie()
	for _, fn := range opts {
		fn(s)
	}
	return s
}

func NewRouter(opts ...Option) Route {
	s := newRouter()
	for _, fn := range opts {
		fn(s)
	}
	return s
}
