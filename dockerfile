FROM golang:1.21.5 AS builder

WORKDIR /app

RUN mkdir -p cmd/stresstest
RUN mkdir -p internal/stresstest/usecase

COPY go.mod ./
COPY cmd/stresstest/main.go ./cmd/stresstest
COPY internal/stresstest/usecase/stresstest_usecase_test.go ./internal/stresstest/usecase
COPY internal/stresstest/usecase/stresstest_usecase.go ./internal/stresstest/usecase

RUN CGO_ENABLED=0 GOOS=linux go build -o go-stress-test cmd/stresstest/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-stress-test .

ENTRYPOINT ["./go-stress-test"]
