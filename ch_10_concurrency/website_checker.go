package ch_10_concurrency

type WebsiteChecker func(string) bool

func CheckWebsite(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = checker(url)
	}

	return results
}
