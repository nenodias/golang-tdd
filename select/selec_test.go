package select_module

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeServerLate(late time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(late)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestCorredor(t *testing.T) {
	t.Run("testing fast url will win the race", func(t *testing.T) {
		serverSlow := makeServerLate(20 * time.Millisecond)
		serverFast := makeServerLate(0 * time.Microsecond)
		defer serverSlow.Close()
		defer serverFast.Close()

		URLSlow := serverSlow.URL
		URLFast := serverFast.URL

		got, _ := Corredor(URLSlow, URLFast)
		want := URLFast

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("return an error if the server won't answer on 10 seconds", func(t *testing.T) {
		server := makeServerLate(25 * time.Millisecond)
		defer server.Close()

		_, err := Configurable(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("an error was expected")
		}
	})
}
