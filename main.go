package main

import "fmt"

// Slice - how to use slices in Go
func Slice() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:5] // Slice từ index 1 đến 3 (không bao gồm index 4)
	fmt.Println("Slice from array:", slice)
}

func Make() {
	value := make([]int, 3, 5)
	value[0] = 1
	value[1] = 2
	value[2] = 3
    fmt.Println("value form make: ", value)
}

func main() {
	fmt.Println("Run program using slice Golang")
	Slice() // Gọi hàm Slice để thực thi ví dụ về slice
	Make()  // Gọi hàm Make để thực thi ví dụ về make
}
