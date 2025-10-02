// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

//counterfeiter:generate -o mocks/metrics_pusher.go --fake-name MetricsPusher . Pusher

// Pusher pushes metrics to a Prometheus Pushgateway.
type Pusher interface {
	// Push pushes all registered metrics to the Pushgateway.
	Push(ctx context.Context) error
	// Gatherer sets the Gatherer to use for collecting metrics.
	Gatherer(gatherer prometheus.Gatherer) Pusher
	// Collector adds a Collector to the Pusher.
	Collector(collector prometheus.Collector) Pusher
	// Client sets a custom HTTP client for the Pusher.
	Client(httpClient push.HTTPDoer) Pusher
}

// NewPusher creates a new Pusher that sends metrics to the specified Pushgateway URL
// with the given job name. Use the fluent API methods to configure collectors, gatherers,
// or a custom HTTP client before calling Push.
func NewPusher(
	url string,
	jobName Name,
) Pusher {
	return &pusher{
		pusher: push.New(
			url,
			jobName.String(),
		),
	}
}

type pusher struct {
	pusher *push.Pusher
}

func (p *pusher) Push(ctx context.Context) error {
	return p.pusher.PushContext(ctx)
}

func (p *pusher) Client(c push.HTTPDoer) Pusher {
	p.pusher.Client(c)
	return p
}

func (p *pusher) Gatherer(g prometheus.Gatherer) Pusher {
	p.pusher.Gatherer(g)
	return p
}

func (p *pusher) Collector(g prometheus.Collector) Pusher {
	p.pusher.Collector(g)
	return p
}
