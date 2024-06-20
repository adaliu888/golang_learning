package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
}

func NewPerson(P string) *person {
	return &person{name: P}
}
func main() {
	//reflect.TypeOf(43).Kind() 调用将返回一个表示 int 类型的 reflect.Kind 常量。在 Go 中，int 类型的 Kind 是 reflect.Int。
	//reflect.TypeOf(43).Name() 调用将返回 int 类型的确切名称，这通常是 int，不过具体名称取决于 int 类型的大小和是否有符号，比如 int、int8、int16、int32 或 int64
	fmt.Println(reflect.TypeOf(43).Kind()) //reflect.Int，类型检查
	fmt.Println(reflect.TypeOf(43).Name()) //
	//reflect.ValueOf(43)

	fmt.Println(reflect.ValueOf(43))        //输出值，获取值
	fmt.Println(reflect.ValueOf(43).Type()) // 输出值的类型,类型断言
	//值转换到原始类型
	fmt.Println(reflect.ValueOf(43).Interface().(int))
	//修改值,使用 reflect.Value 的 CanSet() 方法检查值是否可修改，然后使用 SetInt(), SetFloat(), SetString() 等方法修改值。
	if reflect.ValueOf(43).CanSet() {
		reflect.ValueOf(43).SetInt(44)
	}

	//具体类型
	var i int64 = 42
	fmt.Println(reflect.TypeOf(i).Kind())
	fmt.Println(reflect.TypeOf(i).Name())
	// reflect.String
	var s string = "hello"
	fmt.Println(reflect.TypeOf(s).Kind())
	fmt.Println(reflect.TypeOf(s).Name())

	//reflect.Array
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(array).Kind())
	//fmt.Println(reflect.TypeOf(array).Name())
	//reflect.Slice
	var slice []int = []int{1, 2, 3, 4, 5} //slice type
	SS := slice[3:]
	fmt.Println(reflect.TypeOf(SS).Kind())
	//fmt.Println(reflect.TypeOf(SS).Name())
	//reflect.Map
	var map1 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(reflect.TypeOf(map1).Kind())
	//fmt.Println(reflect.TypeOf(map1).Name())
	//reflect.Struct
	type person struct {
		name string
		age  int
	}
	var p person = person{"John", 30}
	fmt.Println(reflect.TypeOf(p).Kind())
	fmt.Println(reflect.TypeOf(p).Name())

	//reflect.Valueof
	ps := reflect.ValueOf(p)
	fmt.Println(ps.Field(1).Int())
	fmt.Println(ps.FieldByName("name").String())

	//调用结构体的方法
	//ps.MethodByName("John").Call([]reflect.Value{})

	//reflect.Ptr
	var p1 *int = new(int)
	fmt.Println(reflect.TypeOf(p1).Kind())
	//fmt.Println(reflect.TypeOf(p1).Name())

	//reflect.Interface
	var i1 interface{} = 1
	fmt.Println(reflect.TypeOf(i1).Kind())
	fmt.Println(reflect.TypeOf(i1).Name())

	//reflect.Un--一个错误
	var name string = "hello" //un type
	fmt.Println(reflect.TypeOf(name).Kind())
	fmt.Println(reflect.TypeOf(name).Name())
	//reflect.Func
	var f func() = func() { fmt.Println("hello") }
	fmt.Println(reflect.TypeOf(f).Kind())
	fmt.Println(reflect.TypeOf(f).Name())

	//reflect.Channel
	ch := make(chan int)
	fmt.Println(reflect.TypeOf(ch).Kind())
	fmt.Println(reflect.TypeOf(ch).Name())

	//reflect.go func()

}
