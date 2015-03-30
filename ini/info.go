// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ini

import (
	"errors"
	"reflect"
	"strings"
)

type info struct {
	elems    map[string]reflect.Value
	sections map[string]map[string]reflect.Value
}

func scan(v interface{}) (*info, error) {
	ret := &info{
		elems:    map[string]reflect.Value{},
		sections: map[string]map[string]reflect.Value{},
	}

	err := scanValue(v, ret.elems, ret.sections)
	return ret, err
}

// 通过sections是否为nil来判断当前是否已经在sections中
func scanValue(obj interface{}, elems map[string]reflect.Value, sections map[string]map[string]reflect.Value) error {
	v := reflect.ValueOf(obj)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("unmarshal:只接受struct指针")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		name := ft.Name
		tag := ft.Tag.Get("ini")
		if len(tag) > 0 {
			if tag[0] == '-' {
				continue
			}
			tags := strings.SplitN(tag, ",", 2)
			if len(name) > 0 {
				name = tags[0]
			}
			if strings.ToLower(tags[1]) == "section" {
				if sections == nil {
					return errors.New("scanValue:不支持多层嵌套")
				}
				items := map[string]reflect.Value{}
				err := scanValue(v.Field(i).Interface(), items, nil)
				if err != nil {
					return err
				}
				sections[name] = items
			}
		}
		elems[name] = v.Field(i)
	}

	return nil
}
