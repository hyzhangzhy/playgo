package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

/*
n个线程，编号 1 2 3 --- n，打印 1 2 3 ------ maxm，要求按编号顺序，然后到maxm的时候全部退出
*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	chans := make([]chan int, n)
	for i := 0; i < n; i++ {
		chans[i] = make(chan int, 1)
	}
	var wg sync.WaitGroup
	idx := 1
	maxm := 100 //最大值
	wg.Add(n)
	// fmt.Fprintln(out, len(chans))
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				<-chans[id]
				ok := true
				if idx <= maxm {
					fmt.Fprintln(out, id+1, idx)
					idx++
				} else {
					ok = false
				}
				chans[(id+1)%n] <- 1
				if !ok {
					break
				}
			}
		}(i)
	}
	chans[0] <- 1
	wg.Wait()

}
