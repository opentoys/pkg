package cache

import (
	"sync"
	"time"
)

type value struct {
	data any
	time time.Time
	exp  time.Duration
}

type data struct {
	mu     sync.RWMutex
	data   map[string]value
	ticker time.Duration
	exp    time.Duration
}

type Option func(*data)

func WithTicker(t time.Duration) Option {
	return func(d *data) {
		d.ticker = t
	}
}

func WithDefaultExpire(exp time.Duration) Option {
	return func(d *data) {
		d.exp = exp
	}
}

func New(size int, opts ...Option) *data {
	var s = &data{
		data:   make(map[string]value, size),
		ticker: time.Minute,
	}
	for _, v := range opts {
		v(s)
	}
	go s.init()
	return s
}

func (s *data) init() {
	var timer = time.NewTicker(s.ticker)
	for {
		<-timer.C
		func() {
			s.mu.Lock()
			defer s.mu.Unlock()
			for k, v := range s.data {
				if v.exp == 0 {
					continue
				}
				if v.time.Add(v.exp).Before(time.Now()) {
					delete(s.data, k)
				}
			}
		}()
	}
}

func (s *data) Set(k string, v any) {
	s.SetEx(k, v, s.exp)
}

func (s *data) SetEx(k string, v any, ex time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = value{
		data: v,
		time: time.Now(),
		exp:  ex,
	}
}

func (s *data) Get(k string) (v any, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[k]
	if !ok {
		return
	}
	if val.time.Add(val.exp).Before(time.Now()) {
		return nil, false
	}
	return val.data, true
}

func (s *data) GetEx(k string) (v any, exp time.Duration, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[k]
	if !ok {
		return
	}
	exp = time.Since(val.time.Add(val.exp))
	return val.data, exp, true
}
