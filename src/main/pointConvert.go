package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func pointConvertTest() {
	var f float64 = 0.11
	u := Float64bits(f)
	fmt.Println(u)

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
