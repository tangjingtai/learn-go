package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func mutexTest() {
	mutex := sync.Mutex{}
	mutex.Lock()
	//mutex.Lock()
	runtime.Gosched()
	fmt.Println("已获取锁")

	mutex.Unlock()

	a := 1
	var wg sync.WaitGroup

	wg.Add(2)
	go func() { //goroutine1
		defer wg.Done()
		a = a + 1
	}()

	go func() { //goroutine2
		defer wg.Done()
		if a == 1 {
			runtime.Gosched()
			fmt.Println("a==", a)
		}
	}()
	runtime.Gosched()
	wg.Wait()
}
