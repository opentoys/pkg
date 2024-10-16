package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type Request interface {
	R() Request
	BeforeRequest(func(*http.Request) error) Request
	AfterResponse(func(*http.Response) error) Request

	SetHeader(k, v string) Request
	SetHeaders(headers map[string]string) Request
	SetBody(v any) Request
	SetJSON(v any) Request
	SetXML(v any) Request
	SetFormData(data map[string]string) Request
	SetMultipart(v any) Request
	SetResult(v any) Request
	SetQueryParam(k, v string) Request
	SetQueryParams(qs map[string]string) Request

	Do(ctx context.Context, method string, url string) (buf []byte, e error)
	Response(ctx context.Context, method string, url string) (resp *http.Response, e error)
	Get(ctx context.Context, url string) ([]byte, error)
	Post(ctx context.Context, url string) ([]byte, error)
}

type Logger interface {
	Infof(ctx context.Context, msg string, keyvals ...any)
}

type Option func(*client)

func WithClient(hc *http.Client) Option {
	return func(c *client) {
		c.c = hc
	}
}

func WithLogger(log Logger) Option {
	return func(s *client) {
		s.log = log
	}
}

type request struct {
	body   any
	result any
	header http.Header
	qs     url.Values
	base   string
}

type client struct {
	c      *http.Client
	log    Logger
	r      *request
	before []func(*http.Request) error
	after  []func(*http.Response) error
}

func New(ops ...Option) *client {
	s := &client{
		c: &http.Client{},
		r: &request{},
	}
	for _, v := range ops {
		v(s)
	}

	return s
}

func (s *client) BeforeRequest(fn func(*http.Request) error) Request {
	s.before = append(s.before, fn)
	return s
}

func (s *client) AfterResponse(fn func(*http.Response) error) Request {
	s.after = append(s.after, fn)
	return s
}

func (s *client) R() Request {
	var c = &client{
		c:   s.c,
		log: s.log,
		r: &request{
			header: http.Header{},
			qs:     url.Values{},
			base:   s.r.base,
		},
	}

	for k, vs := range s.r.header {
		for _, v := range vs {
			c.r.header.Add(k, v)
		}
	}

	for k, vs := range s.r.qs {
		for _, v := range vs {
			c.r.qs.Add(k, v)
		}
	}
	return c
}

func (s *client) SetHeader(k, v string) Request {
	s.r.header.Add(k, v)
	return s
}

func (s *client) SetHeaders(data map[string]string) Request {
	for k, v := range data {
		s.r.header.Add(k, v)
	}
	return s
}

func (s *client) SetBody(v any) Request {
	s.r.body = v
	return s
}

func (s *client) SetJSON(v any) Request {
	s.r.header.Add("Content-Type", "application/json")
	s.r.body = v
	return s
}

func (s *client) SetXML(v any) Request {
	s.r.header.Add("Content-Type", "text/xml")
	s.r.body = v
	return s
}

func (s *client) SetFormData(data map[string]string) Request {
	s.r.header.Add("Content-Type", "application/x-www-form-urlencoded")
	var qs url.Values
	for k, v := range data {
		qs.Add(k, v)
	}
	s.r.body = qs.Encode()
	return s
}

func (s *client) SetMultipart(v any) Request {
	var buf bytes.Buffer
	var body = multipart.NewWriter(&buf)
	switch vv := v.(type) {
	case map[string]string:
		for k, v := range vv {
			w, _ := body.CreateFormField(k)
			_, _ = w.Write([]byte(v))
		}
	case map[string]io.Reader:
		for k, v := range vv {
			w, _ := body.CreateFormField(k)
			_, _ = io.Copy(w, v)
		}
	case *multipart.Form:
		for k, v := range vv.Value {
			for _, iv := range v {
				w, _ := body.CreateFormField(k)
				_, _ = w.Write([]byte(iv))
			}
		}
		for k, v := range vv.File {
			for _, iv := range v {
				f, e := iv.Open()
				if e != nil {
					continue
				}
				w, _ := body.CreateFormFile(k, iv.Filename)
				_, _ = io.Copy(w, f)
			}
		}
	}
	_ = body.Close()
	s.r.header.Add("Content-Type", body.FormDataContentType())
	s.r.body = &buf
	return s
}

func (s *client) SetResult(v any) Request {
	s.r.result = v
	return s
}

func (s *client) SetQueryParam(k, v string) Request {
	s.r.qs.Add(k, v)
	return s
}

func (s *client) SetQueryParams(data map[string]string) Request {
	for k, v := range data {
		s.r.qs.Add(k, v)
	}
	return s
}

func (s *client) request(ctx context.Context, method string, url string) (req *http.Request, e error) {
	if s.r.base != "" && !strings.HasPrefix(url, "http") {
		url = path.Join(s.r.base, url)
	}
	if s.r.body == nil {
		req, e = http.NewRequestWithContext(ctx, method, url, nil)
		return
	}
	switch vv := s.r.body.(type) {
	case *strings.Reader:
		req, e = http.NewRequestWithContext(ctx, method, url, vv)
	case *bytes.Buffer:
		req, e = http.NewRequestWithContext(ctx, method, url, vv)
	case []byte:
		req, e = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(vv))
	case io.Reader:
		req, e = http.NewRequestWithContext(ctx, method, url, vv)
	case string:
		req, e = http.NewRequestWithContext(ctx, method, url, bytes.NewBufferString(vv))
	default:
		var buf []byte
		if strings.Contains(s.r.header.Get("Content-Type"), "xml") {
			buf, e = xml.Marshal(s.r.body)
			if e != nil {
				return
			}
			req, e = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buf))
			return
		}
		buf, e = json.Marshal(s.r.body)
		if e != nil {
			return
		}
		req, e = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buf))
	}
	return
}

func (s *client) Do(ctx context.Context, method string, url string) (buf []byte, e error) {
	req, e := s.request(ctx, method, url)
	if e != nil {
		return
	}

	for _, fn := range s.before {
		if e = fn(req); e != nil {
			return
		}
	}

	var start = time.Now()
	s.log.Infof(ctx, "http request: %+v", map[string]any{
		"header": req.Header,
		"body":   req.Body,
		"url":    req.URL,
		"method": req.Method,
	})

	var resp *http.Response
	defer func() {
		if resp == nil {
			s.log.Infof(ctx, "http response: %+v", map[string]any{
				"url":    req.URL,
				"method": req.Method,
				"error":  e.Error(),
				"time":   time.Since(start),
			})
			return
		}
		s.log.Infof(ctx, "http response: %+v", map[string]any{
			"url":    req.URL,
			"method": req.Method,
			"header": resp.Header,
			"body":   string(buf),
			"code":   resp.StatusCode,
			"time":   time.Since(start),
		})
	}()

	resp, e = s.c.Do(req)
	if e != nil && errors.Is(e, context.DeadlineExceeded) {
		e = errors.New("请求接口超时")
		return
	} else if e != nil {
		return
	}

	for _, fn := range s.after {
		if e = fn(resp); e != nil {
			return
		}
	}

	buf, e = io.ReadAll(resp.Body)
	return
}

func (s *client) Response(ctx context.Context, method string, url string) (resp *http.Response, e error) {
	req, e := s.request(ctx, method, url)
	if e != nil {
		return
	}

	for _, fn := range s.before {
		if e = fn(req); e != nil {
			return
		}
	}

	var start = time.Now()
	s.log.Infof(ctx, "http request: %+v", map[string]any{
		"header": req.Header,
		"body":   req.Body,
		"url":    req.URL,
		"method": req.Method,
	})

	defer func() {
		if resp == nil {
			s.log.Infof(ctx, "http response: %+v", map[string]any{
				"url":    req.URL,
				"method": req.Method,
				"error":  e.Error(),
				"time":   time.Since(start),
			})
			return
		}
		s.log.Infof(ctx, "http response: %+v", map[string]any{
			"url":    req.URL,
			"method": req.Method,
			"header": resp.Header,
			"code":   resp.StatusCode,
			"time":   time.Since(start),
		})
	}()

	resp, e = s.c.Do(req)
	if e != nil && errors.Is(e, context.DeadlineExceeded) {
		e = errors.New("请求接口超时")
		return
	} else if e != nil {
		return
	}

	for _, fn := range s.after {
		if e = fn(resp); e != nil {
			return
		}
	}

	return
}

func (s *client) Get(ctx context.Context, url string) ([]byte, error) {
	return s.Do(ctx, http.MethodGet, url)
}

func (s *client) Post(ctx context.Context, url string) ([]byte, error) {
	return s.Do(ctx, http.MethodPost, url)
}
