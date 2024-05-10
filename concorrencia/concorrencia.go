package concorrencia

type VerifierWebsite func(string) bool
type Result struct {
	string
	bool
}

func VerifyWebsites(vw VerifierWebsite, urls []string) map[string]bool {
	got := make(map[string]bool)
	channelResult := make(chan Result)
	for _, url := range urls {
		go func(u string) {
			channelResult <- Result{u, vw(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-channelResult
		got[result.string] = result.bool
	}
	return got
}
