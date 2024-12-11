package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in golang")

	myCh := make(chan int, 1)

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(val, isChannelOpen)
		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		// close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
