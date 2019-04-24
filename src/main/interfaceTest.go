package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "dog"
}

func Test(params interface{}) {
	fmt.Println(params)
}

func interfaceTest() {
	animals := []Animal{Cat{}, Dog{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	Test("string")
	Test(123)
	Test(true)
}
