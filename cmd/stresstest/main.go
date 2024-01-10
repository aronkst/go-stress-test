package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aronkst/go-stress-test/internal/stresstest/usecase"
)

func main() {
	flagUrl := flag.String("url", "", "URL of the service to be tested")
	flagRequests := flag.Int("requests", 0, "Total number of requests")
	flagConcurrency := flag.Int("concurrency", 1, "Number of simultaneous calls")
	flag.Parse()

	url := *flagUrl
	requests := *flagRequests
	concurrency := *flagConcurrency

	if url == "" || requests == 0 || concurrency == 0 {
		fmt.Println("Error loading input parameters")
		os.Exit(0)
	}

	report := usecase.ExecuteStressTest(url, requests, concurrency)

	fmt.Println("Load Test Report:")
	fmt.Printf("- Total Time Spent in Milliseconds:   %d\n", int(report.TotalTime/time.Millisecond))
	fmt.Printf("- Total Number of Requests:           %d\n", report.TotalRequests)
	fmt.Printf("- Number of Requests with Status 200: %d\n", report.SuccessfulRequests)

	for code, count := range report.ErrorStatusCodes {
		fmt.Printf("- Status Code %d:                    %d\n", code, count)
	}
}
