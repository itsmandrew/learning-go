package reflection

import "reflect"

// What is interfacee{}?
// You may come across scenarios where you want to writee a function where
// you don't know the type at compile time

// Go lets us go around this with the type interace{}, which is just ANY

// As a writer of a function, you have to be able to inspect anything that has been passed
// to you and try and figure out what you can do with it.
func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	// Can not use NumField on a pointer VCalue
	// Reworking so we check type first and the n do our work
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}

	}

}

// Encapsulating

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
