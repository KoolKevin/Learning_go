package main

import "fmt"

func foo() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}

		close(c1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}

		close(c2)
	}()

	i := 0
	for i < 2 {
		select {
		case val, ok := <-c1:
			if !ok {
				i++
				c1 = nil
			} else {
				fmt.Println("da goroutine 1:", val)
			}

		case val, ok := <-c2:
			if !ok {
				i++
				c2 = nil
			} else {
				fmt.Println("da goroutine 2:", val)
			}
		}
	}
}

func main() {
	foo()
}
