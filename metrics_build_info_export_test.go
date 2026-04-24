// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import "github.com/prometheus/client_golang/prometheus"

// BuildInfoGauge exposes the package-level gauge for tests to inspect
// values via prometheus/testutil. Only used from *_test.go in the
// metrics_test package.
var BuildInfoGauge = buildInfoGauge

// ResetBuildInfoGauge clears all label series on the gauge so tests
// start from a clean slate.
func ResetBuildInfoGauge() {
	buildInfoGauge.Reset()
}

// UnregisterBuildInfoGauge removes the gauge from the default registry.
// Tests should not normally need this; provided for completeness in
// case a future test wants to replace the registration.
func UnregisterBuildInfoGauge() bool {
	return prometheus.DefaultRegisterer.Unregister(buildInfoGauge)
}
