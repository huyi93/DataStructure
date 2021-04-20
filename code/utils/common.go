package utils

import (
	"errors"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 判断是否包含(target 是否包含 obj)
// 支持 string array slice map
func Contain(obj interface{}, target interface{}) (bool, error) {

	kind := reflect.TypeOf(target).Kind()
	value := reflect.ValueOf(target)
	objKind := reflect.TypeOf(obj).Kind()
	objValue := reflect.ValueOf(obj)

	switch kind {
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if value.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	case reflect.String:
		if objKind == reflect.String {
			// 枚举
			for i := 0; i+objValue.Len() < value.Len(); i++ {
				if objValue.String() == value.String()[i:i+objValue.Len()] {
					return true, nil
				}
			}
		}
	}
	return false, errors.New("target not in source")
}
