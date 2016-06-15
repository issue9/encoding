// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"testing"

	"github.com/issue9/assert"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	semver := &SemVersion{}
	a.NotError(Parse(semver, "2.3.19")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19)

	a.NotError(Parse(semver, "2.3.19+build.1")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.Build, "build.1")

	a.NotError(Parse(semver, "2.3.19-pre.release+build")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.PreRelease, "pre.release").
		Equal(semver.Build, "build")

	a.NotError(Parse(semver, "2.3.19-pre.release")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.PreRelease, "pre.release")
}
