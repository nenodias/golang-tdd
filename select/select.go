package select_module

import (
	"fmt"
	"net/http"
	"time"
)

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}

func mesureRequestTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	return time.Since(start)
}

const tenSecondsLimit = 10 * time.Second

func Corredor(a, b string) (string, error) {
	return Configurable(a, b, tenSecondsLimit)
}

func Configurable(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout for %s, and %s", a, b)
	}
}
