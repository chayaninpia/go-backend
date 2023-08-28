package tox

import (
	"reflect"
)

func Struct[T any](s interface{}) T {

	// s is nil
	if s == nil {
		return s.(T)
	}

	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		return Struct[T](reflect.Indirect(reflect.ValueOf(s)).Interface())
	}

	// s is value
	if vx, ok := s.(T); ok {
		return vx
	}

	return s.(T)
}
