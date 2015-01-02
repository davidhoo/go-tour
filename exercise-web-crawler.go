package main

import (
	"fmt"
)

func pad(l int, c string) string {
	l = 4 - l
	s := ""
	for i := 0; i < l; i++ {
		s += c
	}
	return s
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
	var c func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool)
	c = func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			return
		}

		history[url] = true

		ch <- fmt.Sprintf("%sfound %s %q\n **** %v\n", pad(depth, "-"), url, body, urls)
		for _, u := range urls {
			if !history[u] {
				history[u] = true
				c(u, depth-1, fetcher, ch, history)
			} else {
				fmt.Printf(">>>%s is repetitive\n", u)
			}
		}
		return
	}

	c(url, depth, fetcher, ch, history)

	close(ch)
}

func main() {
	ch := make(chan string)
	history := make(map[string]bool)
	go Crawl("http://golang.org/", 4, fetcher, ch, history)
	for s := range ch {
		fmt.Println(s)
	}

	fmt.Println("============= history ================")
	fmt.Println(history)
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
