package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

// go test concorrencia/*.go -race
// Para validar se existem race conditions
func TestVerifyWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := VerifyWebsites(mockVerifierWebsite, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func mockVerifierWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubVerifierWebsite(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkVerifieWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}
	for i := 0; i < b.N; i++ {
		VerifyWebsites(slowStubVerifierWebsite, urls)
	}
}
