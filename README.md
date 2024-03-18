# Go StressTest

This is a Go CLI for performing load tests on a web service.

## Prerequisites

Make sure you have Docker installed on your system. You can download and install Docker from the [official Docker website](https://www.docker.com/).

## Usage with Docker

1. Clone the repository:

```bash
git clone https://github.com/user/project.git
```

2. Navigate to the project directory:

```bash
cd go-stress-test
```

3. Build the Docker image:

```bash
docker build -t go-stress-test .
```

4. Run the CLI using Docker, with the desired parameters:

```bash
docker run --rm go-stress-test --url <SERVICE_URL> --requests <NUMBER_OF_REQUESTS> --concurrency <SIMULTANEOUS_CALLS>
```

Replace `<SERVICE_URL>`, `<NUMBER_OF_REQUESTS>`, and `<SIMULTANEOUS_CALLS>` with the desired values.

## Project Structure

- `cmd/stresstest/main.go`: Contains the main logic of the CLI.
- `internal/stresstest/usecase/stresstest_usecase.go`: Implementation of the use case for performing load tests.
