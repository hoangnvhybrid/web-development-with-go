package main

import "fmt"

func main() {
	//message := make(chan string, 2)
	//message <- "Golang"
	//message <- "Gopher"
	//// nhan gia tri tu kenh
	//fmt.Println(<- message)
	//fmt.Println(<- message)

	message1 := make(chan int)
	go func(ch chan int) {
		fmt.Println(<-ch)
	}(message1)
	message1 <- 1
	//message1 <- 2
	//message1 <- 3
	//fmt.Println(<- message1)
}
