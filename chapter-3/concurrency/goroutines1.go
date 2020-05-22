package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//khai bao so luong goroutines, count = 2, khi mot goroutine goi wg.Done, thi count se giam xuong, khi count giam ve 0 thi cac lenh sau wg.Wait se duoc thuc hien
	wg.Add(2)
	fmt.Println("Start goroutine")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printCounts1(i)
			// o day se nhin thay moi luc la i = 10, vi i duoc in ra truoc khi ham printCounts1 chay xong
			fmt.Println("Finish printCounts1(i): ", i)
		}()
	}
	go printCounts("A")
	go printCounts("B")
	fmt.Println("Waiting to Finish")
	//doi den khi count = 0, thi cac lenh phia sau se duoc chay
	wg.Wait()
	fmt.Println("\nTerminating program")
}

func printCounts1(label int) {
	for count := 100; count < 110; count++ {
		fmt.Printf("Count: %d from %d\n", count, label)
	}
}

func printCounts(label string) {
	//bao cho ham mai goroutine da thuc thi xong
	defer wg.Done()
	for count := 1; count < 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
