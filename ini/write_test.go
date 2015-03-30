// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ini

import (
	"bytes"
	"testing"

	"github.com/issue9/assert"
)

type writeTester struct {
	tokens []*Token
	value  string
}

var testData = []*writeTester{
	// 比较正规的写法
	&writeTester{
		tokens: []*Token{
			&Token{Type: Comment, Value: "comment"},
			&Token{Type: Element, Value: "Value1", Key: "Key1"},
			&Token{Type: Element, Value: "Value2", Key: "Key2"},
			&Token{Type: Section, Value: "section1"},
			&Token{Type: Element, Value: "Value1", Key: "Key1"},
		},
		value: `#comment
Key1=Value1
Key2=Value2
[section1]
Key1=Value1
`,
	},

	// 多行注释
	&writeTester{
		tokens: []*Token{
			&Token{Type: Comment, Value: "comment line 1"},
			&Token{Type: Comment, Value: "comment line 2"},
			&Token{Type: Comment, Value: ""}, // 空行
			&Token{Type: Element, Value: "value", Key: "key"},
		},
		value: `#comment line 1
#comment line 2
#
key=value
`,
	},

	// 带转义字符和注释行空格
	&writeTester{
		tokens: []*Token{
			&Token{Type: Comment, Value: "comment 1\n\n"},
			&Token{Type: Element, Value: "value", Key: "key"},
			&Token{Type: Comment, Value: "\n comment 3 "},
			&Token{Type: Element, Value: "value", Key: "key"},
			&Token{Type: Comment, Value: " \ncomment 4"},
		},
		value: `#comment 1
#
#
key=value
#
# comment 3 
key=value
# 
#comment 4
`,
	},
}

func TestWrite(t *testing.T) {
	a := assert.New(t)
	buf := new(bytes.Buffer)

	for index, test := range testData {
		buf.Reset()
		w := NewWriter(buf, '#')
		a.NotNil(w)
		for _, token := range test.tokens {
			switch token.Type {
			case Comment:
				w.AddComment(token.Value)
			case Element:
				w.AddElementf(token.Key, token.Value)
			case EOF:
				break
			case Section:
				w.AddSection(token.Value)
			case Undefined:
				t.Errorf("在第[%v]个测试数据中检测到Type值为Undefined的Token", index)
			default:
				t.Errorf("在第[%v]个测试数据中检测到Type值[%v]为未定义的Token", index, token.Type)
			}
		} // end for test.tokens
		w.Flush()
		a.Equal(buf.String(), test.value)
	}
}
