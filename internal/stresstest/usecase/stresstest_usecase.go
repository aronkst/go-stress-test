package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Report struct {
	TotalTime          time.Duration
	TotalRequests      int
	SuccessfulRequests int
	ErrorStatusCodes   map[int]int
}

func ExecuteStressTest(url string, requests int, concurrency int) Report {
	var successCount int
	var errorStatusCodes = make(map[int]int)
	var wg sync.WaitGroup

	startTime := time.Now()
	done := make(chan bool)

	requestFunc := func() {
		defer wg.Done()

		client := http.Client{Timeout: time.Second * 10}

		for i := 0; i < requests/concurrency; i++ {
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("Error when making the request:", err)
				continue
			}

			if resp.StatusCode == http.StatusOK {
				successCount++
			} else {
				errorStatusCodes[resp.StatusCode]++
			}

			resp.Body.Close()
		}

		done <- true
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go requestFunc()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {
		// used for synchronization
	}

	elapsedTime := time.Since(startTime)

	report := Report{
		TotalTime:          elapsedTime,
		TotalRequests:      requests,
		SuccessfulRequests: successCount,
		ErrorStatusCodes:   errorStatusCodes,
	}

	return report
}
