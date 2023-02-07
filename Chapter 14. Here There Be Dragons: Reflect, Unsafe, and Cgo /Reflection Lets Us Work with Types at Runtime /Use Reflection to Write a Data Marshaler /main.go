package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type MyData struct {
	Name   string `csv:"name"`
	HasPet bool   `csv:"has_pet"`
	Age    int    `csv:"age"`
}

//	//now to turn entries into output
//	out, err := Marshal(entries)
//	if err != nil {
//		panic(err)
//	}
//	sb := &strings.Builder{}
//	w := csv.NewWriter(sb)
//	w.WriteAll(out)
//	fmt.Println(sb)
//}

func main() {
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"`
	r := csv.NewReader(strings.NewReader(data))
	all, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var entities []MyData
	err = Unmarshal(all, &entities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(entities)

	out, err := Marshal(entities)
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	err = w.WriteAll(out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sb)
}

func Marshal(slice any) ([][]string, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := v.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string

	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < v.Len(); i++ {
		row, err := marshalOne(v.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}

	return out, nil

}

func marshalOne(v reflect.Value) ([]string, error) {
	var out []string
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			out = append(out, strconv.FormatInt(f.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			out = append(out, strconv.FormatUint(f.Uint(), 10))
		case reflect.String:
			out = append(out, f.String())
		case reflect.Bool:
			out = append(out, strconv.FormatBool(f.Bool()))
		default:
			return nil, errors.New("cannot handle field of kind " + f.Kind().String())
		}
	}
	return out, nil
}

func marshalHeader(structType reflect.Type) []string {
	var out []string
	for i := 0; i < structType.NumField(); i++ {
		f := structType.Field(i)
		if tag, ok := f.Tag.Lookup("csv"); ok {
			out = append(out, tag)
		}
	}
	return out
}

func Unmarshal(data [][]string, ptrOfSliceOfStruct any) error {
	sliceValPtr := reflect.ValueOf(ptrOfSliceOfStruct)
	if sliceValPtr.Kind() != reflect.Pointer {
		return errors.New("must be a pointer to a ptrOfSliceOfStruct of structs")
	}
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a ptrOfSliceOfStruct of structs")
	}

	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a ptrOfSliceOfStruct of structs")
	}

	header := data[0]
	namePos := make(map[string]int, structType.NumField())
	for index, column := range header {
		namePos[column] = index
	}
	for _, row := range data[1:] {
		elem := reflect.New(structType).Elem()
		err := unmarshalRow(row, namePos, elem)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, elem))
	}
	return nil
}

func unmarshalRow(row []string, pos map[string]int, elem reflect.Value) error {
	t := elem.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		p, ok := pos[tf.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[p]
		vf := elem.Field(i)
		switch vf.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			vf.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			vf.SetUint(i)
		case reflect.String:
			vf.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			vf.SetBool(b)
		default:
			return errors.New("cannot handle field kind " + vf.Kind().String())
		}

	}
	return nil
}
