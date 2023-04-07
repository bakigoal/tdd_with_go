package ch_10_concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, checker(u)}
		}(url)
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
