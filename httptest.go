package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	total_time := time.Duration(0)
	num_passes := 10
	var url string 

	// Expects the URL on the commandline
	num_args := len(os.Args)
	if num_args >= 2 {
		url = os.Args[1]
	} else {
		fmt.Println("Incluye la URL en la línea de comando")
		os.Exit(1)
	}

	if num_args == 3 {
        	if result, err := strconv.Atoi(os.Args[2]); err == nil {
			num_passes = result
		}	
	}

	// Assigns number of passes if given on CL

	// Prime DNS
	_, err := http.Head(url)
	if err != nil {
		fmt.Println("Había un problema conectando")
		os.Exit(1)
	}

	// Perform the passes, keeping track of how long each takes
	for i := 0; i < num_passes; i++ {
		start := time.Now()
		_, _ = http.Head(url)
		elapsed := time.Since(start)

		fmt.Printf("Pass %d took %s\n", i, elapsed)

		total_time += elapsed
	}

	average := total_time / time.Duration(num_passes)
	fmt.Printf("Total time for %d passes = %s\n", num_passes, total_time)
	fmt.Printf("Average time for each pass = %s\n", average)
}
