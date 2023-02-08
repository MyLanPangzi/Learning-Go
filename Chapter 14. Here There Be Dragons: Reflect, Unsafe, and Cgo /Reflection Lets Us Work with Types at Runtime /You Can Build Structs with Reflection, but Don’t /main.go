package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

func main() {
	f, err := Memorizer(AddSlowly, 2*time.Second)
	if err != nil {
		panic(err)
	}
	addSlowlyCached := f.(func(int, int) int)
	for i := 0; i < 5; i++ {
		now := time.Now()
		r := addSlowlyCached(1, 2)
		fmt.Println("get result", r, "in", time.Since(now))
	}
	time.Sleep(3 * time.Second)
	now := time.Now()
	r := addSlowlyCached(1, 2)
	fmt.Println("get result", r, "in", time.Since(now))

}

func Memorizer(f any, expiration time.Duration) (interface{}, error) {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return nil, errors.New("only for functions")
	}
	inType, err := buildInStruct(ft)
	if err != nil {
		return nil, err
	}
	if ft.NumOut() == 0 {
		return nil, errors.New("must have at least one returned value")
	}
	cache := map[any]outExp{}
	fv := reflect.ValueOf(f)
	return reflect.MakeFunc(ft, func(in []reflect.Value) (out []reflect.Value) {
		p := reflect.New(inType).Elem()
		for i, v := range in {
			p.Field(i).Set(v)
		}
		pv := p.Interface()

		ov, ok := cache[pv]
		now := time.Now()
		if !ok || ov.expiry.Before(now) {
			ov.out = fv.Call(in)
			ov.expiry = now.Add(expiration)
			cache[pv] = ov
		}
		return ov.out
	}).Interface(), nil
}

func buildInStruct(ft reflect.Type) (reflect.Type, error) {
	if ft.NumIn() == 0 {
		return nil, errors.New("must have at least one param")
	}
	fields := make([]reflect.StructField, 0, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		t := ft.In(i)
		if !t.Comparable() {
			return nil, fmt.Errorf("parameter %d of type %s and kind %v is not comparable", i+1, t.Name(), t.Kind())
		}
		fields = append(fields, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: t,
		})
	}
	return reflect.StructOf(fields), nil
}

func AddSlowly(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}
