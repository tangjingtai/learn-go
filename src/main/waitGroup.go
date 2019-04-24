package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupTest() {
	var swg sync.WaitGroup
	for i := 0; i < 1; i++ {
		//增加一个计数器
		swg.Add(1)
		go func(wg *sync.WaitGroup, mark int) {
			//减去计数器
			defer wg.Done() //等价于 wg.Add(-1)
			//if i == 4{
			time.Sleep(time.Second * 2)
			//}
			fmt.Printf("%d goroutine finish \n", mark)
		}(&swg, i)
	}
	//time.Sleep(time.Second * 1)
	//等待所有go程结束
	swg.Wait()
}
