// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics_test

import (
	"time"

	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus/testutil"

	"github.com/bborbe/metrics"
)

var _ = Describe("BuildInfoMetrics", func() {
	BeforeEach(func() {
		metrics.ResetBuildInfoGauge()
	})

	Describe("SetBuildInfo", func() {
		It("sets the gauge to the build timestamp keyed by version and commit", func() {
			buildDate := libtime.DateTime(time.Date(2026, 4, 24, 12, 30, 0, 0, time.UTC))
			bim := metrics.NewBuildInfoMetrics()

			bim.SetBuildInfo("v1.2.3", "abc1234", &buildDate)

			gauge := metrics.BuildInfoGauge.WithLabelValues("v1.2.3", "abc1234")
			Expect(testutil.ToFloat64(gauge)).To(Equal(float64(buildDate.Unix())))
		})

		It("is a no-op when buildDate is nil", func() {
			bim := metrics.NewBuildInfoMetrics()

			bim.SetBuildInfo("v1.2.3", "abc1234", nil)

			// No series should have been created for any label combination.
			Expect(testutil.CollectAndCount(metrics.BuildInfoGauge)).To(Equal(0))
		})

		It("keeps each (version, commit) pair as a distinct time series", func() {
			earlier := libtime.DateTime(time.Date(2026, 4, 20, 10, 0, 0, 0, time.UTC))
			later := libtime.DateTime(time.Date(2026, 4, 24, 12, 30, 0, 0, time.UTC))
			bim := metrics.NewBuildInfoMetrics()

			bim.SetBuildInfo("v1.2.3", "abc1234", &earlier)
			bim.SetBuildInfo("v1.2.4", "def5678", &later)

			Expect(testutil.CollectAndCount(metrics.BuildInfoGauge)).To(Equal(2))
			Expect(
				testutil.ToFloat64(metrics.BuildInfoGauge.WithLabelValues("v1.2.3", "abc1234")),
			).To(Equal(float64(earlier.Unix())))
			Expect(
				testutil.ToFloat64(metrics.BuildInfoGauge.WithLabelValues("v1.2.4", "def5678")),
			).To(Equal(float64(later.Unix())))
		})

		It("overwrites the value when called twice with the same labels", func() {
			earlier := libtime.DateTime(time.Date(2026, 4, 20, 10, 0, 0, 0, time.UTC))
			later := libtime.DateTime(time.Date(2026, 4, 24, 12, 30, 0, 0, time.UTC))
			bim := metrics.NewBuildInfoMetrics()

			bim.SetBuildInfo("v1.2.3", "abc1234", &earlier)
			bim.SetBuildInfo("v1.2.3", "abc1234", &later)

			Expect(testutil.CollectAndCount(metrics.BuildInfoGauge)).To(Equal(1))
			Expect(
				testutil.ToFloat64(metrics.BuildInfoGauge.WithLabelValues("v1.2.3", "abc1234")),
			).To(Equal(float64(later.Unix())))
		})

		It("preserves dirty-version labels as distinct series", func() {
			clean := libtime.DateTime(time.Date(2026, 4, 24, 10, 0, 0, 0, time.UTC))
			dirty := libtime.DateTime(time.Date(2026, 4, 24, 12, 30, 0, 0, time.UTC))
			bim := metrics.NewBuildInfoMetrics()

			bim.SetBuildInfo("v1.2.3", "abc1234", &clean)
			bim.SetBuildInfo("v1.2.3-dirty", "abc1234", &dirty)

			Expect(testutil.CollectAndCount(metrics.BuildInfoGauge)).To(Equal(2))
		})
	})
})
