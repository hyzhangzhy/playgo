package main

import (
	"fmt"
	"sync"
)

/*
3个goroutine 各打印a, b, c

abcabc
*/

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("a")
			ch1 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("b")
			ch2 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("c")
			fmt.Println("------------------------")
			ch3 <- 1
		}
	}()
	ch3 <- 1
	wg.Wait()

}
