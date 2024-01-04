# Web Crawler in Go

This Go program is a simple web crawler that extracts URLs from a webpage up to a specified depth.

## How to Run

1. Make sure you have Go installed on your machine. If not, you can download and install it from [here](https://golang.org/dl/).

2. Clone the repository:

    ```bash
    git clone https://github.com/akhiljns/go-crawler.git
    ```

3. Change into the project directory:

    ```bash
    cd go-crawler
    ```

4. Run the program:

    ```bash
    go run main.go -url https://www.example.com/ -depth 2
    ```

    Replace the URL and depth parameters as needed.

5. **Branch Information:**
   - The master branch contains the concurrent code. To switch to the synchronous version, run:

     ```bash
     git checkout sync-crawler
     ```

6. **Code Progression:**
   - To understand how the code progressed from synchronous to concurrent, please to go through the commit history. Commits provide insights into the step-by-step development process (can take a look at https://github.com/akhiljns/go-crawler/pull/1/files)

## Explanation of the Code

- **Visited Map Usage:**
  - The code utilizes a `visited` map to keep track of visited URLs, preventing redundant processing.
  
- **Initialization with CrawlWebpage:**
  - The `CrawlWebpage` function sets the initial depth for the root URL and initiates the crawling process.
  
- **Recursive Crawling with Crawl Function:**
  - The `crawl` function fetches HTML content, extracts links, and spawns goroutines for unvisited URLs.
  
- **Asynchronous Processing with Goroutines:**
  - Goroutines are used for asynchronous processing, enhancing performance by parallelizing URL crawling.
  
- **Concurrency Safety with Mutex:**
  - To ensure data consistency, a mutex (`visitedMutex`) is employed to synchronize access to the `visited` map.
  
- **Efficient and Safe Crawling:**
  - The combination of recursion, goroutines, and mutex provides an efficient and safe mechanism for web crawling in a concurrent environment.

## Code Assessment

1. **Correctness:**
   - The program correctly crawls web pages and extracts URLs. Ensure you have a stable internet connection and correct URL input.

2. **Simplicity:**
   - The code is straightforward, following a standard web crawling pattern using goroutines and a wait group.

3. **Maintainability:**
   - The code is reasonably maintainable; functions are relatively short and focused on specific tasks.

4. **Performance:**
   - The performance is acceptable for a simple web crawler. Consider adding rate limiting for web requests for improved politeness towards servers.

