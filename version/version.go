// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 版本号解析工具
package version

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// 表示结构体字段的类型，版本号要嘛是字符串，要嘛是数值
const (
	fieldTypeNumber = iota
	fieldTypeString
)

// 对每个字段的描述
type field struct {
	Type  int
	Seq   map[byte]int
	Value reflect.Value
}

// 解析版本号字符串到 obj 中。
func Parse(obj interface{}, ver string) error {
	fields, err := getFields(obj)
	if err != nil {
		return err
	}

	start := 0
	field := fields[0]
	for i := 0; i < len(ver)+1; i++ {
		var nextIndex int
		var found bool

		if i < len(ver) { // 未结束字符串
			b := ver[i]
			nextIndex, found = field.Seq[b]
			if !found {
				continue
			}
		}

		switch field.Type {
		case fieldTypeNumber:
			n, err := strconv.ParseInt(ver[start:i], 10, 64)
			if err != nil {
				return err
			}
			field.Value.SetInt(n)
		case fieldTypeString:
			field.Value.SetString(ver[start:i])
		default:
			return errors.New("未知道的 fieldType" + strconv.Itoa(field.Type))
		}

		i++ // 过滤掉当前字符
		start = i
		field = fields[nextIndex] // 下一个 field
	} // end for

	return nil
}

func getFields(obj interface{}) ([]*field, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("参数 obj 只能是结构体")
	}
	t := v.Type()

	fields := make([]*field, v.NumField(), v.NumField())
	for i := 0; i < v.NumField(); i++ {
		tags := strings.Split(t.Field(i).Tag.Get("version"), ",")
		if len(tags) < 2 {
			return nil, errors.New("缺少必要的标签元素")
		}

		index, err := strconv.Atoi(tags[0])
		if err != nil {
			return nil, err
		}

		field := &field{Seq: make(map[byte]int, 2)}
		switch tags[1] {
		case "number":
			field.Type = fieldTypeNumber
		case "string":
			field.Type = fieldTypeString
		default:
			return nil, errors.New("无效的标签：" + tags[1])
		}

		for _, v := range tags[2:] {
			n, err := strconv.Atoi(v[1:])
			if err != nil {
				return nil, err
			}
			field.Seq[v[0]] = n
		}

		field.Value = v.Field(i)

		fields[index] = field
	}

	return fields, nil
}
