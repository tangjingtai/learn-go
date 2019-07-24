package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

type P struct {
	Id   int
	Name string
}

func pointConvertTest() {
	var f float64 = 0.11
	u := Float64bits(f)
	fmt.Println(u)
	//(*[2]unsafe.Pointer)(unsafe.Pointer(i))[1]
	p := &P{1, "zhangsan"}
	ps := (*[2]unsafe.Pointer)(unsafe.Pointer(p))
	p0 := ps[0]
	p1 := ps[1]
	fmt.Println(p, ps, p0, p1)

	us := [...]uint32{1, 2, 3}
	p_us := unsafe.Pointer(&us)
	u64 := (*uint64)(p_us)
	fmt.Println(u64, *u64)

	fmt.Println("地址", p_us, &us[0], &us[1], &us[2])
	//var u2 uint= 100
	//p:= &u2
	//up := uintptr(unsafe.Pointer(p))
	//p2 := unsafe.Pointer(up + 2)
	//u3:=*(*uint)(p2)
	//fmt.Println(p2,u3)
}
