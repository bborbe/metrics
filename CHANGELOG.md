# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v0.5.1

- Add unit tests for `BuildInfoMetrics` covering value set, nil no-op, distinct label series, overwrite, and `-dirty` version support

## v0.5.0

- Add `BuildInfoMetrics` interface for recording build provenance as Prometheus metric
- Add `build_info` GaugeVec with version/commit/timestamp labels
- Add `github.com/bborbe/time` dependency
- Bump `golang.org/x/vuln` to v1.3.0 and other stdlib deps

## v0.4.8

- Bump golangci-lint to v2.11.4
- Bump osv-scanner to v2.3.5
- Bump gosec to v2.25.0
- Update multiple indirect dependencies
- Add new OSV scanner vulnerability ignores

## v0.4.7

- Update numerous indirect dependencies (docker, containerd, opentelemetry, golang.org/x/*)
- Replace k8s.io/kube-openapi replace directive with charmbracelet/x/cellbuf, denis-tingaikin/go-header, opencontainers/runtime-spec
- Remove large exclude block for k8s.io and sigs.k8s.io packages
- Add new indirect deps: clipperhouse/displaywidth, clipperhouse/stringish, clipperhouse/uax29

## v0.4.6

- chore: verified project health — precommit passes with exit code 0, no issues found

## v0.4.5

- chore: verified project health — all tests pass, linting clean, precommit succeeds

## v0.4.4

- upgrade golangci-lint from v1 to v2
- standardize Makefile: add .PHONY declarations, multiline trivy, mocks mkdir
- update .golangci.yml to v2 format
- setup dark-factory config

## v0.4.3

- go mod update

## v0.4.2

- go mod update

## v0.4.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v0.4.0

- update go and deps

## v0.3.3
- Add package-level documentation (doc.go) with examples
- Enhance Pusher interface documentation with fluent API pattern explanation
- Update README example to demonstrate proper context propagation pattern

## v0.3.2
- Fix project name in Makefile (corrected from skeleton to metrics)

## v0.3.1
- Rename parameter from 'name' to 'jobName' in NewPusher for better clarity

## v0.3.0

- fix testsmak

## v0.2.0

- Add fluent API to Pusher with Gatherer, Collector, and Client methods
- Add comprehensive godoc comments to all exported types and functions
- Add test suite for Pusher with mock Pushgateway server
- Update README with complete usage examples and fluent API documentation
- Remove default gatherer from NewPusher, requiring explicit configuration

## v0.1.0

- Initial release with Name builder and Pusher
