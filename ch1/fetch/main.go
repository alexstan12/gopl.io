// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://"){
			//url = "http://" + url
			var builder strings.Builder
			builder.WriteString("http://")
			builder.WriteString(url)
			url = builder.String()
			fmt.Println(url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("HTTP status code: %v", resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}

	// for _,url := range os.Args[1:] {
	// 	response, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v", err)
	// 		os.Exit(1)
	// 	}
	// 	contents, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "body reading: %v", err)
	// 		os.Exit(1)
	// 	}
	// 	response.Body.Close()
	// 	fmt.Printf("%s", string(contents))
	// }
}

//!-
