package main

import (
	"fmt"
	"reflect"
)

func ValueOfSet() {
	var i int = 1024
	valueOfi := reflect.ValueOf(&i)
	valueOfi = valueOfi.Elem()

	valueOfi2 := reflect.ValueOf(i)
	fmt.Println("valueOfi2.CanAddr(): ", valueOfi2.CanAddr())
	fmt.Println("valueOfi2.CanSet(): ", valueOfi2.CanSet())

	valueOfi.SetInt(2048)
	fmt.Println("valueOfi: ", valueOfi.Int())
	fmt.Println("i: ", i)

	structFieldSet()
}

// 结构体实例可寻址，结构体中的字段也可寻址
func structFieldSet() {
	type dog struct {
		LegCount int
	}
	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()
	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	vLegCount.SetInt(4)
	fmt.Println("LegCount", vLegCount.Int())
}

func reflectNewInstance() {
	var i int = 10
	typeOfi := reflect.TypeOf(i)
	instance := reflect.New(typeOfi)

	fmt.Printf("instance.Type():%s,instance.Kind():%s\n", instance.Type(), instance.Kind())
	fmt.Println("instance.Elem().CanSet()： ", instance.Elem().CanSet())
	fmt.Println("instance.Elem().CanAddr()： ", instance.Elem().CanAddr())

	instance.Elem().SetInt(100)
	fmt.Println("i:", instance.Elem().Int())
}

func reflectCall() {
	add := func(a, b int) int {
		return a + b
	}

	valueOfAdd := reflect.ValueOf(add)
	valueOfResult := valueOfAdd.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200)})
	fmt.Println("result:", valueOfResult[0].Int())
}
