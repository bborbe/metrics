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
    "github.com/prometheus/client_golang/prometheus"
)

// Create a registry and register your metrics
prometheusRegistry := prometheus.NewRegistry()

// Create and configure a pusher using the fluent API
pusher := metrics.NewPusher(
    "http://monitoring-pushgateway:9091",
    metrics.BuildName("kafka", "backup", "my_topic"),
).Gatherer(prometheusRegistry)

// Push metrics to the gateway
if err := pusher.Push(context.Background()); err != nil {
    // handle error
}
```

The pusher supports a fluent API for configuration:

```go
pusher := metrics.NewPusher(url, jobName).
    Gatherer(registry).              // Use a custom registry
    Collector(myCollector).          // Add a specific collector
    Client(customHTTPClient)         // Use a custom HTTP client
```
