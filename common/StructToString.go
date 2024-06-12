package common

import (
	"fmt"
	"reflect"
)

func StructToString(v interface{}) string {
	rv := reflect.ValueOf(v)
	rt := rv.Type()
	var result string

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		result += fmt.Sprintf("%s: %v, ", rt.Field(i).Name, field.Interface())
	}

	if len(result) > 2 {
		result = result[:len(result)-2] // 移除最后一个逗号和空格
	}

	return result
}
