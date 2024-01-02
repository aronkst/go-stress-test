package usecase_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aronkst/go-stress-test/internal/stresstest/usecase"
)

func TestExecuteStressTest(t *testing.T) {
	type args struct {
		requests            int
		concurrency         int
		httpHeader          int
		httpTimeout         int
		wantSuccessRequests int
		wantRequestsError   int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "successful stress test",
			args: args{
				requests:            10,
				concurrency:         5,
				httpHeader:          http.StatusOK,
				httpTimeout:         0,
				wantSuccessRequests: 10,
				wantRequestsError:   0,
			},
		},
		{
			name: "error stress test",
			args: args{
				requests:            10,
				concurrency:         5,
				httpHeader:          http.StatusInternalServerError,
				httpTimeout:         0,
				wantSuccessRequests: 0,
				wantRequestsError:   10,
			},
		},
		{
			name: "timeout stress test",
			args: args{
				requests:            10,
				concurrency:         5,
				httpHeader:          http.StatusOK,
				httpTimeout:         11,
				wantSuccessRequests: 0,
				wantRequestsError:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(time.Duration(tt.args.httpTimeout) * time.Second)
				w.WriteHeader(tt.args.httpHeader)
			}))
			defer testServer.Close()

			startTime := time.Now()
			report := usecase.ExecuteStressTest(testServer.URL, tt.args.requests, tt.args.concurrency)
			elapsedTime := time.Since(startTime)

			if report.TotalTime > elapsedTime {
				t.Errorf("[%s] TotalTime = %d; want <= %d", tt.name, report.TotalTime, elapsedTime)
			}

			if report.TotalRequests != tt.args.requests {
				t.Errorf("[%s] TotalRequests = %d; want %d", tt.name, report.TotalRequests, tt.args.requests)
			}

			if report.SuccessfulRequests != tt.args.wantSuccessRequests {
				t.Errorf("[%s] SuccessfulRequests = %d; want %d", tt.name, report.SuccessfulRequests, tt.args.wantSuccessRequests)
			}

			if report.ErrorStatusCodes[http.StatusInternalServerError] != tt.args.wantRequestsError {
				t.Errorf("[%s] report.ErrorStatusCodes[http.StatusInternalServerError] = %d; want %d", tt.name, report.ErrorStatusCodes[http.StatusInternalServerError], tt.args.wantRequestsError)
			}
		})
	}
}
