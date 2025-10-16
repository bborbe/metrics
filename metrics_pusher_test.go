// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/bborbe/metrics"
)

var _ = Describe("Pusher", func() {
	var server *ghttp.Server
	var ctx context.Context
	var pusher metrics.Pusher
	var registry *prometheus.Registry

	BeforeEach(func() {
		ctx = context.Background()
		server = ghttp.NewServer()
		registry = prometheus.NewRegistry()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Push", func() {
		It("pushes metrics successfully", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("PUT", "/metrics/job/test_job"),
					ghttp.RespondWith(200, ""),
				),
			)

			pusher = metrics.NewPusher(
				server.URL(),
				metrics.BuildName("test", "job"),
			).Gatherer(registry)

			err := pusher.Push(ctx)
			Expect(err).To(BeNil())
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})

		It("returns error on failed push", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("PUT", "/metrics/job/test_job"),
					ghttp.RespondWith(500, ""),
				),
			)

			pusher = metrics.NewPusher(
				server.URL(),
				metrics.BuildName("test", "job"),
			).Gatherer(registry)

			err := pusher.Push(ctx)
			Expect(err).NotTo(BeNil())
		})
	})
})
