package tox

import (
	"reflect"
)

func Slice[T any](s interface{}) []T {

	// s is nil
	if s == nil {
		return nil
	}

	if reflect.ValueOf(s).IsZero() {
		return nil
	}

	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		return Slice[T](reflect.Indirect(reflect.ValueOf(s)).Interface())
	}

	// s is value
	if vx, ok := s.([]T); ok {
		return vx
	}

	return nil
}
