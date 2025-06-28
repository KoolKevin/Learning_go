package main

import (
	"fmt"
	"sync"
)

func foo() {
	var c chan int = make(chan int)
	var done chan bool = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for val := range c {
			fmt.Println(val)
		}

		done <- true
	}()

	wg.Wait()
	close(c)

	<-done
}

func main() {
	foo()
}
