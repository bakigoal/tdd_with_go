package ch_12_reflection

import "reflect"

func Walk(x interface{}, f func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), f)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			f(field.String())
		case reflect.Struct:
			Walk(field.Interface(), f)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
