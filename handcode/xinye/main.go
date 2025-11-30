package main

import (
	//"reflect"
	"fmt"
	"sync"
	//"sync/atomic"
)

/*
写一段高效的多协程代码，有3个协程同时运行，每个线程依次从0打印到9，希望输出结果如下：
000111222333444555666777888999
*/

func main() {
	wg := sync.WaitGroup{}
	cha := make(chan int, 1)
	chb := make(chan int)
	chc := make(chan int)
	//count := int32(0)
	a := func() {
		defer wg.Done()
		for i := 0; i < 900; i++ {
			<-cha
			fmt.Println(i)
			chb <- 1

		}
	}
	b := func() {
		defer wg.Done()
		for i := 0; i < 900; i++ {
			<-chb
			fmt.Println(i)
			chc <- 1

		}
	}
	c := func() {
		defer wg.Done()
		for i := 0; i < 900; i++ {
			<-chc
			fmt.Println(i)
			cha <- 1

		}
	}

	wg.Add(3)
	cha <- 1
	go a()
	go b()
	go c()
	wg.Wait()

}
