package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Variables with command line flags
var url string
var numPasses int
var ignoreStatus bool
var hostname string

func init() {
	flag.StringVar(&hostname, "h", "", "Hostname for Host header")
	flag.IntVar(&numPasses, "n", 10, "Number of connections")
	flag.BoolVar(&ignoreStatus, "s", false, "Ignore returned HTTP status code")
	flag.StringVar(&url, "u", "", "URL, including protocol (HTTP/HTTPS), mandatory")
}

func main() {
	var totalTime time.Duration = 0

	flag.Parse()

	//Check if url is empty, if so exit
	if url == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Prepare the HTTP client with URL and HEADERS
	client := &http.Client{}
	req, _ := http.NewRequest("HEAD", url, nil)

	if hostname != "" {
		req.Host = hostname
	}

	// Prime DNS and verify connectivity
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 && !ignoreStatus {
		log.Fatal("Received a non-200 status code: ", resp.StatusCode)
	}

	// Perform the passes, keeping track of how long each takes
	for i := 0; i < numPasses; i++ {
		start := time.Now()
		_, err := client.Do(req)
		elapsed := time.Since(start)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pass %d took %s\n", i, elapsed)
		totalTime += elapsed

	}

	// Compute average and print stats
	average := totalTime / time.Duration(numPasses)
	fmt.Printf("Total time for %d passes = %v\n", numPasses, totalTime)
	fmt.Printf("Average time for each pass = %v\n", average)
}
