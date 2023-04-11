package ch_12_reflection

import "reflect"

func Walk(x interface{}, f func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		f(field.String())
	}
}
