package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//nhu vay la khi minh ko chu dong sleep thi goroutine se chay xong roi goroutines khac moi duoc chay, con neu minh chu dong sleep thi goroutine nay dang chay se bi dung de goroutine khac chay tiep theo
	//khai bao so luong goroutines, count = 2, khi mot goroutine goi wg.Done, thi count se giam xuong, khi count giam ve 0 thi cac lenh sau wg.Wait se duoc thuc hien
	wg.Add(2)
	fmt.Println("Start goroutine")
	go printCounts("A")
	go printCounts("B")
	fmt.Println("Waiting to Finish")
	//doi den khi count = 0, thi cac lenh phia sau se duoc chay
	wg.Wait()
	fmt.Println("\nTerminating program")
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
