package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Print(prefix string, ok bool) {
	fmt.Printf("%s: Name is %s, Age is %d :%t\n", prefix, u.Name, u.Age, ok)
}

func main() {
	u := User{"张三", 20}

	//TypeOf
	t := reflect.TypeOf(u)
	fmt.Println(t)

	//Kind
	fmt.Println(t.Kind())

	//ValueOf
	v := reflect.ValueOf(u)
	fmt.Println(v)

	//%T %v
	//fmt.Printf("%T\n", u)
	//fmt.Printf("%v\n", u)

	//将v转成u
	u1 := v.Interface().(User)
	fmt.Println(u1.Name, u1.Age)

	//遍历字段名和方法
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	//修改变量的值
	vp := reflect.ValueOf(&u) //必须传入变量地址
	vp.Elem().Field(0).SetString("李四")
	fmt.Println(u)

	//调用方法
	mPrint := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀"), reflect.ValueOf(true)} //参数是[]reflect.Value, 通过reflect.ValueOf转
	fmt.Println(mPrint.Call(args))
}
