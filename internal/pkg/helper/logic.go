package helper

import (
	"fmt"
	"log"
	"reflect"
)

// ValueIf 实现 3 元运算符逻辑： logic ? v1 : v2
func ValueIf[T any](logic bool, v1 T, v2 T) T {
	if logic {
		return v1
	}
	return v2
}

// IsNil 是否为 nil
// 支持 chan, func, interface, map, pointer, or slice
func IsNil(v interface{}) bool {
	vi := reflect.ValueOf(v)
	switch vi.Kind() {
	case reflect.Chan:
	case reflect.Ptr:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Func:
	case reflect.Slice:
		return vi.IsNil()
	}
	return true
}

// Must 没有错误，如果出现错误抛出异常
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s:%v", msg, err))
	}
	log.Println("Success:", msg)
}
