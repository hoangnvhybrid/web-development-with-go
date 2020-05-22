package main

import (
	"fmt"
	"sync"
)

//func sumArray(myArr []int, n int, wg *sync.WaitGroup) int {
//	defer wg.Done()
func sumArray(myArr []int, n int) int {
	fmt.Println("Goroutines: ", n)
	sum := 0
	for _, value := range myArr {
		sum += value
	}

	fmt.Println("sum: ", sum)
	return sum
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	myArr1 := []int{2, 3}
	myArr2 := []int{7, 8}
	total := 0
	//go func() {
	//	total += sumArray(myArr1, 1, &wg)
	//}()
	//go func() {
	//	total += sumArray(myArr2, 2, &wg)
	//}()
	go func() {
		defer wg.Done()
		total += sumArray(myArr1, 1)
	}()
	go func() {
		defer wg.Done()
		total += sumArray(myArr2, 2)
	}()
	//time.Sleep(2 * time.Millisecond)
	wg.Wait()
	fmt.Println("total: ", total)
}
