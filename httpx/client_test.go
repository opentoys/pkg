package httpx_test

import (
	"context"
	"easyadmin/internal/resources/logger"
	"easyadmin/internal/valobj"
	"easyadmin/pkg/github/httpx"
	"fmt"
	"sync"
	"testing"
)

func TestXxx(t *testing.T) {
	log := logger.New("default")
	log.Config = &valobj.ConfigLogger{Level: -1}
	log.Register()
	c := httpx.New(httpx.WithLogger(log))
	resp, e := c.R().Get(context.Background(), "https://mail.163.com/")
	fmt.Println(resp, e)
}

func TestBatch(t *testing.T) {
	log := logger.New("default")
	log.Config = &valobj.ConfigLogger{Level: -1}
	log.Register()
	var wg sync.WaitGroup
	c := httpx.New(httpx.WithLogger(log))
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			c := c.R()
			resp, e := c.Get(context.Background(), "https://pay.myscrm.cn")
			fmt.Println(resp, e)
		}()
	}
	wg.Wait()

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			c := c.R()
			resp, e := c.Get(context.Background(), "https://pay.myscrm.cn")
			fmt.Println(resp, e)
		}()
	}
	wg.Wait()
}
