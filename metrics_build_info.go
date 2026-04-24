// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	libtime "github.com/bborbe/time"
	"github.com/prometheus/client_golang/prometheus"
)

var buildInfoGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "build",
		Name:      "info",
		Help:      "Build metadata. Value is build timestamp (Unix seconds). Labels identify the release. Service identified by Prometheus job label.",
	},
	[]string{"version", "commit"},
)

func init() {
	prometheus.MustRegister(buildInfoGauge)
}

//counterfeiter:generate -o mocks/build-info-metrics.go --fake-name BuildInfoMetrics . BuildInfoMetrics

// BuildInfoMetrics records build-time provenance as a Prometheus metric.
//
// Emits a GaugeVec "build_info{version, commit}" whose value is the build
// timestamp in Unix seconds. Service identification comes from the
// Prometheus "job" label set by the scrape config — not a metric label —
// so one "build_info" metric is shared across every binary that imports
// this package.
//
// Typical queries:
//
//	time() - build_info          // process age in seconds
//	count by (version) (build_info)   // distinct versions currently running
type BuildInfoMetrics interface {
	// SetBuildInfo records build provenance for the current process.
	// Call exactly once at startup, after argument parsing.
	// Does nothing when buildDate is nil (local `go run` without build args).
	SetBuildInfo(version, commit string, buildDate *libtime.DateTime)
}

// NewBuildInfoMetrics returns a BuildInfoMetrics writing to the package-level
// "build_info" gauge. The gauge is registered against prometheus.DefaultRegisterer
// at package init.
func NewBuildInfoMetrics() BuildInfoMetrics {
	return &buildInfoMetrics{}
}

type buildInfoMetrics struct{}

// SetBuildInfo sets the gauge to the build timestamp keyed by version+commit.
// No-op when buildDate is nil.
func (m *buildInfoMetrics) SetBuildInfo(
	version, commit string,
	buildDate *libtime.DateTime,
) {
	if buildDate == nil {
		return
	}
	buildInfoGauge.WithLabelValues(version, commit).Set(float64(buildDate.Unix()))
}
