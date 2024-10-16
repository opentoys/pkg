package cache_test

import (
	"easyadmin/pkg/github/cache"
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	m := cache.New(512)
	m.Set("hello", "world")
	m.SetEx("1111", "sakjhdajksdhaksd", time.Second)

	fmt.Println(m.Get("1111"))

	<-time.After(time.Second * 5)

	fmt.Println(m.Get("1111"))

	<-time.After(time.Minute)

	fmt.Println(m.Get("1111"))
}
