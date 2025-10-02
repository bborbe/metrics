# Metrics

Go library for working with Prometheus metrics.

## Features

- **Name Builder**: Create valid Prometheus metric names from arbitrary strings
- **Pusher**: Push metrics to Prometheus Pushgateway

## Installation

```bash
go get github.com/bborbe/metrics
```

## Usage

### Building Metric Names

```go
import "github.com/bborbe/metrics"

// Create a valid Prometheus metric name
name := metrics.BuildName("my-service", "request_count")
// Result: "my_service_request_count"

// Add to existing name
name = name.Add("total")
// Result: "my_service_request_count_total"
```

The `BuildName` function:
- Converts to lowercase
- Replaces illegal characters with underscores
- Handles leading numbers
- Collapses multiple underscores

### Pushing Metrics

```go
import (
    "context"
    "github.com/bborbe/metrics"
)

pusher := metrics.NewPusher(
    "http://pushgateway:9091",
    metrics.BuildName("my_service"),
)

if err := pusher.Push(context.Background()); err != nil {
    // handle error
}
```

## Development

```bash
make precommit  # Format, generate, test, and check
make test       # Run tests
make check      # Run linters and security checks
```

## License

BSD-style license. Copyright Benjamin Borbe.
