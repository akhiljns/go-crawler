package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// CrawlWebpage crawls the given rootURL looking for <a href=""> tags
// that are targeting the current web page, either via an absolute url like http://mysite.com/mypath or by a relative url like /mypath
// and returns a sorted list of absolute urls  (eg: []string{"http://mysite.com/1","http://mysite.com/2"})

// CrawlWebpage crawls a webpage up to a specified depth and returns a list of discovered URLs

func CrawlWebpage(rootURL string, maxDepth int) ([]string, error) {
	var (
		visited      = make(map[string]int)
		visitedMutex sync.Mutex
		wg           sync.WaitGroup
	)

	visited[rootURL] = 0 // Set the initial depth for the root URL or seed

	wg.Add(1)
	go func(url string, depth int) {
		defer wg.Done()
		crawl(maxDepth, url, depth, visited, &visitedMutex, &wg)
	}(rootURL, 0)

	wg.Wait()

	var result []string
	for url := range visited {
		result = append(result, url)
	}

	if len(result) == 0 || len(result) == 1 {
		return nil, nil
	}

	sort.Strings(result)

	return result, nil
}

func crawl(maxDepth int, url string, depth int, visited map[string]int, visitedMutex *sync.Mutex, wg *sync.WaitGroup) {
	if depth >= maxDepth {
		return
	}

	body, err := fetchHTML(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	links := extractLinks(body)

	visitedMutex.Lock()
	for _, link := range links {
		absoluteURL, err := resolveURL(url, link)
		if err != nil {
			fmt.Printf("Error resolving URL %s: %v\n", link, err)
			continue
		}

		if _, exists := visited[absoluteURL]; !exists {
			visited[absoluteURL] = depth + 1 // Update depth based on the parent's depth
			wg.Add(1)
			go func(url string, depth int) {
				defer wg.Done()
				crawl(maxDepth, url, depth, visited, visitedMutex, wg)
			}(absoluteURL, depth+1)
		}
	}
	visitedMutex.Unlock()
}

func resolveURL(base, target string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %v", err)
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		return "", fmt.Errorf("failed to parse target URL: %v", err)
	}

	absoluteURL := baseURL.ResolveReference(targetURL)
	return absoluteURL.String(), nil
}

// fetchHTML makes an HTTP GET request to the given URL and returns the response body
func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// extractLinks parses HTML content and returns a slice of unique URLs found in anchor tags
func extractLinks(htmlContent string) []string {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

// --- DO NOT MODIFY BELOW ---

func main() {
	const (
		defaultURL      = "https://www.example.com/"
		defaultMaxDepth = 2
	)
	urlFlag := flag.String("url", defaultURL, "the url that you want to crawl")
	maxDepth := flag.Int("depth", defaultMaxDepth, "the maximum number of links deep to traverse")
	flag.Parse()

	links, err := CrawlWebpage(*urlFlag, *maxDepth)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	fmt.Println("Links")
	fmt.Println("-----")
	for i, l := range links {
		fmt.Printf("%03d. %s\n", i+1, l)
	}
	fmt.Println()
}
