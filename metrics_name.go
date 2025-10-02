// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"regexp"
	"strings"
)

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) Add(name string) Name {
	return BuildName(n.String(), name)
}

var replaceLeadingNumber = regexp.MustCompile(`^[0-9]`)

var replaceIllegalCharacters = regexp.MustCompile(`[^a-z0-9]+`)

var replaceMultiUnderscore = regexp.MustCompile(`_+`)

// BuildName from the given string. Replace all illegal characters with underscore
func BuildName(names ...string) Name {
	name := strings.Join(names, "_")
	name = strings.ToLower(name)
	name = replaceLeadingNumber.ReplaceAllString(name, "_")
	name = replaceIllegalCharacters.ReplaceAllString(name, "_")
	name = replaceMultiUnderscore.ReplaceAllString(name, "_")
	return Name(name)
}
