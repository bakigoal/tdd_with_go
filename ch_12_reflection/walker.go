package ch_12_reflection

import "reflect"

func Walk(x interface{}, f func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			f(field.String())
		}

		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), f)
		}
	}
}
