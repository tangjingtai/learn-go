package main

import "fmt"

func main() {
	//goroutineEntrance()
	contextTest()
}

func sliceTest() {
	arr := []int{10, 20, 30, 40, 50}
	arr2 := arr[1:3]
	fmt.Println("arr2:", arr2)
	arr2[0] = 2
	fmt.Println("arr:", arr)
	arr2 = append(arr2, 4)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr:", arr)

}
