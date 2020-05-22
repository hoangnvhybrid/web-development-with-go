package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg1 sync.WaitGroup

func main() {
	count := make(chan int)
	// Add a count of two, one for each goroutine.
	wg1.Add(2)
	fmt.Println("Start Goroutines")

	go printCounts2("B", count)
	go printCounts2("A", count)
	go printCounts2("H", count)
	go printCounts2("C", count)
	go printCounts2("D", count)

	fmt.Println("Channel begin")
	count <- 1
	fmt.Println("Waiting To Finish")
	wg1.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts2(label string, count chan int) {
	defer wg1.Done()
	for {
		//Receives message from Channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 100 {
			fmt.Printf("Channel Closed from %s \n", label)
			//close the channel
			close(count)
			return
		}
		val++
		//send count back to the other goroutine
		count <- val
	}
}
