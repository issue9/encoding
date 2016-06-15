// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"testing"

	"github.com/issue9/assert"
)

func TestSemVersion_Compare(t *testing.T) {
	a := assert.New(t)

	v1 := &SemVersion{
		Major: 1,
	}
	v2 := &SemVersion{
		Major: 1,
	}

	a.True(v1.Compare(v2) == 0)

	v1.Minor = 2
	a.True(v1.Compare(v2) > 0)
	v2.Minor = 2

	v2.Patch = 3
	a.True(v1.Compare(v2) < 0)
	v1.Patch = 3

	// build 不参与运算
	v1.Build = "111"
	v2.Build = "222"
	a.True(v1.Compare(v2) == 0)

	v1.PreRelease = "alpha"
	a.True(v1.Compare(v2) < 0)
	a.True(v2.Compare(v1) > 0)

	v2.PreRelease = "beta"
	a.True(v1.Compare(v2) < 0)

	// 相等的 preRelease
	v1.PreRelease = "1.alpha"
	v2.PreRelease = "1.alpha"
	a.True(v1.Compare(v2) == 0)

	// preRelease 数值的比较
	v1.PreRelease = "11.alpha"
	v2.PreRelease = "9.alpha"
	a.True(v1.Compare(v2) > 0)
}

func TestSemVersion_String(t *testing.T) {
	a := assert.New(t)

	sv := &SemVersion{
		Major: 1,
	}
	a.Equal(sv.String(), "1.0.0")

	sv.Minor = 22
	a.Equal(sv.String(), "1.22.0")

	sv.Patch = 1234
	a.Equal(sv.String(), "1.22.1234")

	sv.Build = "20160615"
	a.Equal(sv.String(), "1.22.1234+20160615")

	sv.PreRelease = "alpha1.0"
	a.Equal(sv.String(), "1.22.1234-alpha1.0+20160615")

	sv.Build = ""
	a.Equal(sv.String(), "1.22.1234-alpha1.0")
}

func TestSemVerCompare(t *testing.T) {
	a := assert.New(t)

	v, err := SemVerCompare("1.0.0", "1.0.0")
	a.NotError(err).True(v == 0)

	v, err = SemVerCompare("1.2.0", "1.0.0")
	a.NotError(err).True(v > 0)

	v, err = SemVerCompare("1.2.0", "1.2.1")
	a.NotError(err).True(v < 0)
}
