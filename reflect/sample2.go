package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	field1 int    `comment`
	Field2 string `comment`
	field3 MyStruct2
}

//	リフレクションを用いて呼び出す場合はメソッド名の先頭を大文字にしてExportする必要がある
func (this MyStruct) Test(name string) {
	fmt.Println("name", name)
}
func (this *MyStruct) TestP(name string) {
	fmt.Println("name", name)
}

type MyStruct2 struct {
	field int
}

func main() {
	ms := MyStruct{78, "sample", MyStruct2{3}}
	v := reflect.ValueOf(ms)
	fmt.Println("v", v)
	fmt.Println("v.Field(0)", v.Field(0).Int())
	fmt.Println("v.Field(1)", v.Field(1))
	fmt.Println("v.Field(2)", v.Field(2))
	//	fmt.Println("v.Field(3)", v.Field(3))
	fmt.Println(`v.FieldByName("field1")`, v.FieldByName("field1"))
	fmt.Println("v.FieldByIndex([]int{1, 0})", v.FieldByIndex([]int{2, 0}))
	fmt.Println("v.FieldByNameFunc", v.FieldByNameFunc(func(name string) bool {
		return name == "field2"
	}))
	fmt.Println("v.NumField()", v.NumField())
	fmt.Println("v.NumMethod()", v.NumMethod())
	fmt.Println("v.NumMethod()", v.MethodByName("Test").Call([]reflect.Value{reflect.ValueOf("sample")}))
	vp := reflect.ValueOf(&ms)
	fmt.Println("v.NumMethod()", vp.MethodByName("TestP").Call([]reflect.Value{reflect.ValueOf("sampleP")}))
	vp.Elem().FieldByName("Field2").SetString("new")
	fmt.Println(ms)
}
