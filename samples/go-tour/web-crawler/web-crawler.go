package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	cache.mux.Lock()
	if _, ok := cache.m[url]; ok {
		cache.mux.Unlock()
		fmt.Printf("%v already fetched\n", url)
		return
	}

	cache.m[url] = errors.New("In Progress")
	cache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)

	cache.mux.Lock()
	cache.m[url] = err
	cache.mux.Unlock()

	if err != nil {
		fmt.Printf("Error for %v -> %v\n", url, err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	done := make(chan bool)
	for i, _url := range urls {
		fmt.Printf("Crawling %v/%v of %v: %v.\n", i, len(urls), url, _url)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(_url)
	}

	for i, _url := range urls {
		fmt.Printf("[%v] %v/%v Waiting for %v.\n", url, i, len(urls), _url)
		<-done
	}

	fmt.Printf("Crawled %v!!!\n", url)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

var cache = struct {
	m   map[string]error
	mux sync.Mutex
}{m: make(map[string]error)}
