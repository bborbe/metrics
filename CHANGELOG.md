# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

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
