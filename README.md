# Go Assessment : WebPage Link Crawler

You are tasked with created a web page link crawler program using the [Go Programming Language](https://go.dev/).
The application is given a `-url` and `-depth` flag and prints the list of Links returned by `CrawlWebpage`.

## Directions

Using this initial repository and the [Go Programming Language](https://go.dev/),

- Crawl the web page starting at the given `-url` for any links within the current website.
  - Only consider links that are for the current website under the given `-url`. 
  - Ignore any outbound links off-site (such as to `http://othersite.com`)
  - Ignore any links to anchor tags like `<a href="#foo">`
  - Ignore links to any protocol other than `http/https` (eg: `ftp`, `mailto`, etc...)
  - Resolve relative links to their absolute paths from the page you are crawling (See: [W3C HTML Spec](https://www.w3.org/TR/PR-html40-971107/struct/links.html#h-12.4.1) and related RFCs [RFC1808](https://datatracker.ietf.org/doc/html/rfc1808), and [RFC2068](https://datatracker.ietf.org/doc/html/rfc2068#section-14.11) for more info)

      *Hint*: the `net/url` package should help a lot
    
      **Relative Link Crash Course**

      Given: You are currently crawling on a page `http://example.org/a/b/c/index.html` (called the `BASE` in the HTML Spec)
      - `<a href="http://example.org/foo">` = `http://example.org/foo`
        - Absolute Link, so use this verbatim
      - `<a href="/bar>` = `http://example.org/bar`
        - This replaces the whole path part of the current `BASE`
      - `<a href="baz">` = `http://example.org/a/b/c/baz` 
        - This is relative to the current `BASE` document, since we are in `/a/b/c` viewing `index.html`, we replace `index.html` with `baz`
      - `<a href="./baz">` = `http://example.org/a/b/c/baz`
        - This is just like above, except it explicitly says its relative to the current folder `.`
      - `<a href="../baz">` = `http://example.org/a/b/baz`
        - This is like the above, except it says we should go one folder up, `/a/b/c` becomes `/a/b` and so the link is `/a/b/baz`
- Follow links found on the page to discover any new links.
- Stop following links deeper than `-depth` levels.

  **Example**

   if the max depth is 2, and the main url is `http://example.com/` and there is a chain of links from `/` -> `/foo` -> `/bar` -> `/baz` -> `/quux`
    - depth 0: crawl `/` and discover the link to `/foo`
    - depth 1: crawl `/foo` and discover `/bar`
    - depth 2: crawl `/bar` and discover `/baz`
    - stop at depth 2, do not crawl `/baz` and therefor will not discover the link to `/quux`
    - Collect the links with their base url to produce a list:
      1. `http://example.com/`
      2. `http://example.com/foo`
      3. `http://example.com/bar`
      4. `http://example.com/baz`

Your solution should be implemented or called by the body of: 

```go
func CrawlWebpage(rootURL string, maxDepth int) ([]string, error) {
	//TODO: Implement Solution
	return nil, errors.New("solution not implemented")
}
```

However, you are encouraged to create any additional functions, structs, or packages you deem appropriate.

## Assessment Criteria

1. Correctness - Does the program produce the right outputs for the given inputs.
2. Simplicity - Does the program solve the problem without unnecessary complexity?
3. Maintainability - How easy it is to read, understand, and change?
4. Performance - Is the solution implemented optimally.
5. Style - Is the code idiomatic Go? Was care given to code style such as names of variables, functions, types, etc.? See: [Google Go Style-Guide](https://google.github.io/styleguide/go/)
