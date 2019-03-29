package main

import (
	"github.com/pkg/profile"
	"time"
)

func joinSlice() []string {
	var arr []string
	for i := 0; i < 100000; i++ {
		// 故意造成多次的切片添加(append)操作, 由于每次操作可能会有内存重新分配和移动, 性能较低
		arr = append(arr, "arr")
	}
	return arr
}

func genProfile() {
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer stopper.Stop()
	joinSlice()
	time.Sleep(1000)
}
