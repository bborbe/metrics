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

// Pusher pushes metrics to a push gateway.
type Pusher interface {
	Push(ctx context.Context) error
}

func NewPusher(
	url string,
	name Name,
) Pusher {
	return &pusher{
		pusher: push.New(
			url,
			name.String(),
		).Gatherer(prometheus.DefaultGatherer),
	}
}

type pusher struct {
	pusher *push.Pusher
}

func (p *pusher) Push(ctx context.Context) error {
	return p.pusher.PushContext(ctx)
}
