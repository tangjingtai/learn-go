package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {

}

func main() {
	typeOfA := reflect.TypeOf(MyStruct{})
	fmt.Printf("type:%s,kind:%s",typeOfA.Name(),typeOfA.Kind())
}