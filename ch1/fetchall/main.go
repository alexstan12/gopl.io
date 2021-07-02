// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// func main() {
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch) // start a goroutine
// 	}
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch) // receive from channel ch
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) // send to channel ch
// 		return
// 	}

// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	resp.Body.Close() // don't leak resources
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
// }

func main(){
	logsFile, err :=os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	var builder strings.Builder
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open, err: %v", err)
	}
	start:= time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		if !strings.HasPrefix(url, "http://"){
				var builder strings.Builder
				builder.WriteString("http://")
				builder.WriteString(url)
				url = builder.String()
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		//fmt.Println(<-ch) // receive from channel ch
		builder.WriteString(<-ch)
		builder.WriteString("\n")
	}
	ending := fmt.Sprintf("Time elapsed: %.2fs\n", time.Since(start).Seconds())
	builder.WriteString(ending)
	result := builder.String()
	_, err = logsFile.WriteString(result)
	if err != nil {
		fmt.Printf("Something occured while writing to file, %v \n", err)
	}
}

func fetch(url string, ch chan string){
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Failed to get response, %v", err)
		return 
	}
	noBytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Failed to copy number of bytes, %v", err)
		return
	}
	finish := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", finish, noBytes, url) // send data to channel ch
}

//!-
