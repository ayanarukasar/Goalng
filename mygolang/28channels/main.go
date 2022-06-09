package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-mych
		fmt.Println(val)
		fmt.Println(isChanelOpen)

		fmt.Println(<-mych)
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 5 //adding a value in the channel
		mych <- 6
		close(mych) //syntax for closing
		// mych <- 5   //adding a value in the channel
		// mych <- 6
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
