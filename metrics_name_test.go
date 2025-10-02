// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics_test

import (
	metrics "bitbucket.apps.seibert-media.net/oc/lib-metrics"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildName", func() {
	It("does not modify valid names", func() {
		Expect(metrics.BuildName("my_value").String()).To(Equal("my_value"))
	})
	It("replace multi underscores", func() {
		Expect(metrics.BuildName("my__value").String()).To(Equal("my_value"))
	})
	It("replace invalid characters", func() {
		Expect(metrics.BuildName("my-value").String()).To(Equal("my_value"))
	})
	It("converts to lower case", func() {
		Expect(metrics.BuildName("MY_VALUE").String()).To(Equal("my_value"))
	})
	It("contains multiplace names", func() {
		Expect(metrics.BuildName("a", "b", "c").String()).To(Equal("a_b_c"))
	})
	It("keeps numbers", func() {
		Expect(metrics.BuildName("a1").String()).To(Equal("a1"))
	})
	It("remove leading zero", func() {
		Expect(metrics.BuildName("1a").String()).To(Equal("_a"))
	})
})
