package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	sanitizedUrl := strings.TrimPrefix(url, "http://")
	sanitizedUrl = strings.TrimPrefix(sanitizedUrl, "https://")
	sanitizedUrl = strings.ReplaceAll(sanitizedUrl, "/", "-")
	filename := fmt.Sprintf("%s-%s.txt", sanitizedUrl, time.Now().Format("15:04:05"))
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("while opening/creating %s: %v", filename, err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
