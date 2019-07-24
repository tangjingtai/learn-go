package main

import (
	"golang.org/x/protobuf/proto"
	"io/ioutil"
	"main/proto"
	"os"
)

func protobufTest() {
	p1 := &test.Person{
		Id:   1,
		Name: "张三",
		Phones: []*test.Phone{
			&test.Phone{Type: test.PhoneType_HOME, Number: "13566666666"},
			&test.Phone{Type: test.PhoneType_WORK, Number: "13566666666"},
		},
		Id2: -2,
	}
	p2 := &test.Person{
		Id:   1,
		Name: "李四",
		Phones: []*test.Phone{
			&test.Phone{Type: test.PhoneType_HOME, Number: "15566666666"},
			&test.Phone{Type: test.PhoneType_WORK, Number: "15566666666"},
		},
		Id2: -2,
	}
	book := &test.ContactBook{
		Persons: []*test.Person{p1, p2},
	}
	data, _ := proto.Marshal(book)
	ioutil.WriteFile("/data/textProto.txt", data, os.ModePerm)
}
