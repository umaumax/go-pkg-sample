package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
}

func RefKind(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Int, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Struct:
		return "struct"
	case reflect.Interface:
		return "interface"
	case reflect.Ptr:
		return "Ptr"
	case reflect.UnsafePointer:
		return "UnsafePointer"
	case reflect.Uintptr:
		return "Uintptr"
	case reflect.Invalid:
		return "Invalid"
	}
	return v.Kind().String()
}

type temp struct {
	name string
}

func p(a ...string) {
	for _, v := range a {
		for _, vv := range v {
			fmt.Print(string(vv))
		}
		fmt.Print("____________________")
	}
	fmt.Println()
}

func sample() {
	//	reflect/value.go
	//	func ValueOf(i interface{}) Value {}
	//	引数にintreface{}があるのでreflect.ValueOf()に値を渡す際にinterface{}に包んでも同じ
	//	map[***]interface{},[]interface{}などについてよく成り立つ
	p("1", "int", RefKind(reflect.ValueOf(1)))
	p("1.0", "float", RefKind(reflect.ValueOf(1.0)))
	p("\"1\"", "string", RefKind(reflect.ValueOf("1")))
	p("interface{}\"sample\"", "string", RefKind(reflect.ValueOf(interface{}("sample"))))
	p("errors.New(\"error\")", "Ptr", RefKind(reflect.ValueOf(errors.New("error"))))
	p("temp{}", "struct", RefKind(reflect.ValueOf(temp{})))
	p("&temp{}", "Ptr", RefKind(reflect.ValueOf(&temp{})))
	m := make(map[string]interface{})
	p("make(map[string]interface{})", "map", RefKind(reflect.ValueOf(m)))
	p("&make(map[string]interface{})", "Ptr", RefKind(reflect.ValueOf(&m)))
	p("new(map[string]interface{})", "Ptr", RefKind(reflect.ValueOf(new(map[string]interface{}))))
	m["struct"] = temp{}
	p("map[string]interface{} temp{}", "struct", RefKind(reflect.ValueOf(m["struct"])))
	p("reflect.ValueOf(123).Interface()", "int", RefKind(reflect.ValueOf(reflect.ValueOf(123).Interface())))
	p("reflect.ValueOf(123)", "struct", RefKind(reflect.ValueOf(reflect.ValueOf(123))))
	p("reflect.ValueOf(&temp{})", "struct", RefKind(reflect.ValueOf(reflect.ValueOf(&temp{}))))
	p("reflect.ValueOf(temp{})", "struct", RefKind(reflect.ValueOf(reflect.ValueOf(temp{}))))
	p("reflect.ValueOf(nil)", "Invalid", RefKind(reflect.ValueOf(nil)))
	p("make([]interface{}, 123)", "slice", RefKind(reflect.ValueOf(make([]interface{}, 123))))
	p("[]interface{}{123, 123.0}", "slice", RefKind(reflect.ValueOf([]interface{}{123, 123.0})))
	p("[...]interface{}{123, 123.0}", "array", RefKind(reflect.ValueOf([...]interface{}{123, 123.0})))
	array := make([]interface{}, 123)
	array[0] = reflect.ValueOf(123).Interface()
	array[1] = reflect.ValueOf(&temp{}).Interface()
	array[2] = reflect.ValueOf(temp{})
	//	reflect.Valueへ変換した後にIndex()だとインタフェース型
	p("array.Index(0) reflect.ValueOf(123).Interface()", "interface", RefKind(reflect.ValueOf(array).Index(0)))
	p("array.Index(1) reflect.ValueOf(&temp{}).Interface()", "interface", RefKind(reflect.ValueOf(array).Index(1)))
	p("reflect.ValueOf(array.Index(1)).Interface() reflect.ValueOf(&temp{}).Interface()", "Ptr", RefKind(reflect.ValueOf(reflect.ValueOf(array).Index(1).Interface())))
	p("array.Index(1) reflect.ValueOf(&temp{}).Interface()", "struct", RefKind(reflect.ValueOf(reflect.ValueOf(array).Index(1))))
	p("reflect.ValueOf(array.Index(2)) reflect.ValueOf(temp{})", "struct", RefKind(reflect.ValueOf(reflect.ValueOf(array).Index(2))))
	p("array.Index(0).Interface()", "int", RefKind(reflect.ValueOf(reflect.ValueOf(array).Index(0).Interface())))
	p(RefKind(rarray().Index(0)))
	p(RefKind(reflect.ValueOf(rarray().Index(0).Interface())))
	p(RefKind(rarray()))
	p(RefKind(rarray().Index(0)))
	return
}

func rarray() reflect.Value {
	array := make([]interface{}, 123)
	array[0] = reflect.ValueOf(123).Interface()
	return reflect.ValueOf(array)
}
