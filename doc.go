// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metrics provides utilities for working with Prometheus metrics.
//
// It includes a name builder for creating valid Prometheus metric names from
// arbitrary strings, and a pusher for sending metrics to a Prometheus
// Pushgateway.
//
// # Name Builder
//
// The BuildName function creates valid Prometheus metric names by joining
// strings with underscores, converting to lowercase, and replacing illegal
// characters:
//
//	name := metrics.BuildName("my-service", "request_count")
//	// Result: "my_service_request_count"
//
//	name = name.Add("total")
//	// Result: "my_service_request_count_total"
//
// # Metrics Pusher
//
// The Pusher interface provides a fluent API for pushing metrics to a
// Prometheus Pushgateway:
//
//	func pushMetrics(ctx context.Context) error {
//	    registry := prometheus.NewRegistry()
//	    // ... register your metrics ...
//
//	    pusher := metrics.NewPusher(url, jobName).
//	        Gatherer(registry).
//	        Push(ctx)
//
//	    return pusher.Push(ctx)
//	}
package metrics
