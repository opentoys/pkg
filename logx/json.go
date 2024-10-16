package logx

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"runtime"
	"time"
)

type jsonhandler struct {
	Out    io.Writer
	Err    io.Writer
	Option *slog.HandlerOptions
	attrs  []slog.Attr
	group  string
	encode func(w io.Writer, v any) error
}

var _ slog.Handler = &jsonhandler{}

type Option func(*jsonhandler)

func WithErrorWriter(w io.Writer) Option {
	return func(h *jsonhandler) {
		h.Err = w
	}
}

func WithLevel(level slog.Level) Option {
	return func(h *jsonhandler) {
		h.Option.Level = level
	}
}

func WithAddSource(add bool) Option {
	return func(h *jsonhandler) {
		h.Option.AddSource = add
	}
}

func WithReplaceAttr(fn func(groups []string, a slog.Attr) slog.Attr) Option {
	return func(h *jsonhandler) {
		h.Option.ReplaceAttr = fn
	}
}

func WithEncode(fn func(w io.Writer, v any) error) Option {
	return func(j *jsonhandler) {
		j.encode = fn
	}
}

func New(o io.Writer, opts ...Option) *jsonhandler {
	if o == nil {
		o = io.Discard
	}
	var s jsonhandler
	s.Option = &slog.HandlerOptions{Level: slog.LevelDebug}
	s.Out = o
	s.encode = s.defaultencode
	for _, v := range opts {
		v(&s)
	}
	if s.Err == nil {
		s.Err = s.Out
	}
	return &s
}

func (s *jsonhandler) clone() *jsonhandler {
	return &jsonhandler{
		Out:    s.Out,
		Err:    s.Err,
		Option: s.Option,
		attrs:  s.attrs,
		group:  s.group,
		encode: s.encode,
	}
}

func (s *jsonhandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= s.Option.Level.Level()
}

func (s *jsonhandler) Handle(ctx context.Context, r slog.Record) (e error) {
	if !s.Enabled(ctx, r.Level) {
		return
	}
	var msg = map[string]any{
		"msg":   r.Message,
		"time":  r.Time.Format(time.RFC3339),
		"level": r.Level.String(),
	}
	if s.group != "" {
		r.AddAttrs(slog.String("group", s.group))
	}

	if s.Option.AddSource || r.Level == slog.LevelError {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		r.AddAttrs(slog.Any(slog.SourceKey, &slog.Source{
			Function: f.Function,
			File:     f.File,
			Line:     f.Line,
		}))
	}

	r.Attrs(func(v slog.Attr) bool { msg[v.Key] = v.Value.Any(); return false })

	for _, v := range s.attrs {
		msg[v.Key] = v.Value.Any()
	}

	var w = s.Out
	if r.Level == slog.LevelError && s.Err != nil {
		w = s.Err
	}

	e = s.encode(w, msg)
	return
}

func (s *jsonhandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	a := s.clone()
	a.attrs = append(a.attrs, attrs...)
	return a
}

func (s *jsonhandler) WithGroup(name string) slog.Handler {
	c := s.clone()
	c.group = name
	return c
}

func (s *jsonhandler) defaultencode(w io.Writer, v any) (e error) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	e = enc.Encode(v)
	return
}
