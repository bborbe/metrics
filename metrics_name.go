// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"regexp"
	"strings"
)

// Name represents a valid Prometheus metric name.
type Name string

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// Add appends the given name to the current Name, creating a new valid Prometheus metric name.
func (n Name) Add(name string) Name {
	return BuildName(n.String(), name)
}

var replaceLeadingNumber = regexp.MustCompile(`^[0-9]`)

var replaceIllegalCharacters = regexp.MustCompile(`[^a-z0-9]+`)

var replaceMultiUnderscore = regexp.MustCompile(`_+`)

// BuildName creates a valid Prometheus metric name from the given strings.
// It joins the strings with underscores, converts to lowercase, replaces
// leading numbers and illegal characters with underscores, and collapses
// multiple consecutive underscores into one.
func BuildName(names ...string) Name {
	name := strings.Join(names, "_")
	name = strings.ToLower(name)
	name = replaceLeadingNumber.ReplaceAllString(name, "_")
	name = replaceIllegalCharacters.ReplaceAllString(name, "_")
	name = replaceMultiUnderscore.ReplaceAllString(name, "_")
	return Name(name)
}
