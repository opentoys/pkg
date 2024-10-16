package injects

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AfterConfig Step = 1
	AfterStore  Step = 50
	AfterServer Step = 100
)

type Config struct {
	Every func(*Object)
	Name  string
	Step  Step
}

type Step int8

var Service = NewInjects()

func NewInjects() *global {
	return &global{
		g: &Graph{logger: log{}},
	}
}

type object struct {
	o    *Object
	step Step
}

type global struct {
	g        *Graph
	data     []*object
	everyfns []func(*Object)
	wg       sync.RWMutex
}

func (s *global) SetDebug(t bool) {
	s.g.debug = t
}

func (s *global) Register(v interface{}, cfgs ...*Config) {
	s.wg.Lock()
	defer s.wg.Unlock()
	if len(cfgs) == 0 {
		cfgs = append(cfgs, &Config{})
	}
	// 第一个名字是别名 Object{}.Name
	var item = &Object{Value: v}
	item.Name = cfgs[0].Name
	if cfgs[0].Every != nil {
		s.everyfns = append(s.everyfns, cfgs[0].Every)
	}
	s.data = append(s.data, &object{o: item, step: cfgs[0].Step})
}

func (s *global) Action() (e error) {
	s.wg.RLock()
	defer s.wg.RUnlock()
	for idx := range s.data {
		if e = s.g.Provide(s.data[idx].o); e != nil {
			return
		}
	}

	if e = s.g.Populate(); e != nil {
		return
	}

	var sb strings.Builder
	sb.WriteString("action inject register func\n")
	sort.Slice(s.data, func(i, j int) bool {
		return s.data[i].step < s.data[j].step
	})
	var arr = make([]string, 0, len(s.data))
	var by int8 = -1
	for idx := range s.data {
		if s.data[idx].o.Value == nil {
			continue
		}
		if s := int8(s.data[idx].step); by != s {
			sb.WriteString(fmt.Sprintf("    [step %d]: %v\n", by, arr))
			arr = []string{}
			by = s
		}
		arr = append(arr, s.data[idx].o.String())
		if v, ok := s.data[idx].o.Value.(interface{ Register() (e error) }); ok {
			if e = v.Register(); e != nil {
				return
			}
		}
	}
	sb.WriteString(fmt.Sprintf("    run at %d: %v\n", by, arr))
	if s.g.debug {
		s.g.logger.Debugf(sb.String())
	}
	return
}

func (s *global) FindObjects(prefix string) (lst []interface{}) {
	s.wg.RLock()
	defer s.wg.RUnlock()
	for idx := range s.data {
		if strings.HasPrefix(s.data[idx].o.Name, prefix) {
			lst = append(lst, s.data[idx].o.Value)
		}
	}
	return
}

type log struct{}

func (log) Debugf(format string, v ...interface{}) {
	format += "\n"
	fmt.Printf(format, v...)
}
