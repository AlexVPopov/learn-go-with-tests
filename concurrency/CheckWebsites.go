package concurrency

// WebsiteChecker takes a website address and reports, wether the website is online or not
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites takes a list of websites and checkes their status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
