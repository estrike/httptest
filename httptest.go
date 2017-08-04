package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"log"
)

func main() {
	totalTime := time.Duration(0)
	numPasses := 10
	var url string 

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

	// Assigns number of passes if given on CL

	// Prime DNS
	_, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the passes, keeping track of how long each takes
	for i := 0; i < numPasses; i++ {
		start := time.Now()
		_, _ = http.Head(url)
		elapsed := time.Since(start)

		fmt.Printf("Pass %d took %s\n", i, elapsed)

		totalTime += elapsed
	}

	average := totalTime / time.Duration(numPasses)
	fmt.Printf("Total time for %d passes = %s\n", numPasses, totalTime)
	fmt.Printf("Average time for each pass = %s\n", average)
}
