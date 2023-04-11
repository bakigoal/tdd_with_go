package ch_12_reflection

import "reflect"

func Walk(x interface{}, f func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	f(field.String())
}
