package main

import "fmt"

func chanTest() {
	msg := make(chan string)
	go func() {
		msg <- "hello"
	}()

	s := <-msg
	fmt.Println(s)
}
