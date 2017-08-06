package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"log"
//	"flag"
)

var totalTime time.Duration = 0
var average time.Duration
var numPasses int = 10
var url string

func init() {
	// Expects the URL on the commandline
	numArgs := len(os.Args)
	if numArgs >= 2 {
		url = os.Args[1]
	} else {
		log.Fatal("URL is needed")
	}

	if numArgs == 3 {
        	if result, err := strconv.Atoi(os.Args[2]); err == nil {
			numPasses = result
		}	
	}
}

func main() {
	// Prepare the HTTP client with URL and HEADERS
	client := &http.Client{}
	req, _ := http.NewRequest("HEAD", url, nil)
	req.Host = "www.example.com"

	// Prime DNS
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal ("Received a non-200 status code: ", resp.StatusCode)
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

	average = totalTime / time.Duration(numPasses)
	fmt.Printf("Total time for %d passes = %v\n", numPasses, totalTime)
	fmt.Printf("Average time for each pass = %v\n", average)
}
