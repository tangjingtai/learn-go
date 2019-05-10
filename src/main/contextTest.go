package main

import (
	"context"
	"fmt"
	"time"
)

func contextTest() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	ret := 0
	for {
		if ret > 0 {
			break
		}
		select {
		case a, b := <-ctx.Done():
			fmt.Println("发生超时,", a, b)
			ret = 1
			break
		default:
			fmt.Println("未超时")
			time.Sleep(time.Millisecond * 500)
		}
	}
	fmt.Println("完成")
}
