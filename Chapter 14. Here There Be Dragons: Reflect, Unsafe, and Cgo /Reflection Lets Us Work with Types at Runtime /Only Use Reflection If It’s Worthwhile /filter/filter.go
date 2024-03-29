package filter

import (
	"reflect"
)

func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		val := sv.Index(i)
		returns := fv.Call([]reflect.Value{val})
		if returns[0].Bool() {
			out = reflect.Append(out, val)
		}
	}
	return out.Interface()
}

func FilterString(s []string, f func(string) bool) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func FilterInt(s []int, f func(int) bool) []int {
	out := make([]int, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
