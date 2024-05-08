package gormplus

import (
	"reflect"
)

// IsZero 判断是否为零值
func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsZero()
}
