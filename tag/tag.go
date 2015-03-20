// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// tag包实现对特定格式的struct tag字符串的分析。
// 并不具有很强的通用性。支持以下两种格式的字符串。
//
// 1. 以分号分隔的字符串，每个子串又以逗号分隔，
// 第一个字符串为键名，之后的字符串组成的数组为键值。如：
//  "id,1;unique;fun,add,1,2;"
//  // 以下将会被解析成：
//  [
//       "id"    :["1"],
//       "unique":nil,
//       "fun"   :["add","1","2"]
//  ]
//
// 2.以分号分隔的字符串，每个子串括号前的字符串为健名，
// 括号中的字符串以逗号分隔组成数组为键值。如：
//  "id(1);unique;fun(add,1,2)"
//  // 以下将会被解析成：
//  [
//       "id"    :["1"],
//       "unique":nil,
//       "fun"   :["add","1","2"]
//  ]
package tag

import (
	"strings"
)

// 当前库的版本
const Version = "0.1.1.150319"

// 将第二种风格的struct tag转换成第一种风格的。
var styleReplace = strings.NewReplacer("(", ",", ")", "")

// 分析tag的内容，并以map的形式返回
func Parse(tag string) map[string][]string {
	ret := make(map[string][]string)

	if len(tag) == 0 {
		return nil
	}

	if strings.IndexByte(tag, '(') > -1 {
		tag = styleReplace.Replace(tag)
	}

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		part = strings.Trim(part, ",")
		items := strings.Split(part, ",")
		ret[items[0]] = items[1:]
	}

	return ret
}

// 从tag中查找名称为name的内容。
// 第二个参数用于判断该项是否存在。
func Get(tag, name string) ([]string, bool) {
	if len(tag) == 0 {
		return nil, false
	}

	if strings.IndexByte(tag, '(') > -1 {
		tag = styleReplace.Replace(tag)
	}

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		part = strings.Trim(part, ",")
		items := strings.Split(part, ",")
		if items[0] == name {
			return items[1:], true
		}
	}

	return nil, false
}

// 功能同Get()函数，但在无法找到的情况下，会返回defVal做为默认值。
func MustGet(tag, name string, defVal ...string) []string {
	if ret, found := Get(tag, name); found {
		return ret
	}

	return defVal
}

// 查询指定名称的项是否存在，若只是查找是否存在该项，
// 使用Has()会比Get()要快上许多。
func Has(tag, name string) bool {
	if len(tag) == 0 {
		return false
	}

	if strings.IndexByte(tag, '(') > -1 {
		tag = styleReplace.Replace(tag)
	}

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		part = strings.Trim(part, ",")
		items := strings.SplitN(part, ",", 2)
		if items[0] == name {
			return true
		}
	}

	return false
}
