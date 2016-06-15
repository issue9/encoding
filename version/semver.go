// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

// http://semver.org/lang/zh-CN/
type SemVersion struct {
	Major      int    `version:"0,number,.1"`
	Minor      int    `version:"1,number,.2"`
	Patch      int    `version:"2,number,+4,-3"`
	PreRelease string `version:"3,string,+4"`
	Build      string `version:"4,string"`
}
