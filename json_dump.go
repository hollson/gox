// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"encoding/json"
)

// 将对象转换为Json字符串
func Json(o interface{}) string {
	dump, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(dump)
}

// 将对象转换为美化的Json字符串
func PrettyJson(o interface{}) string {
	dump, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(dump)
}
