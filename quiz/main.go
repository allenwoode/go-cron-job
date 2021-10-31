package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int64)

	go func() {
		for {
			ch <- time.Now().Unix()
			time.Sleep(time.Second)
		}
	}()

	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	t1 := time.NewTimer(20*time.Second)
	defer t1.Stop()

	for {
		select {
		case <-t1.C:
			// timeout
			fmt.Println("timeout 20s")
		case <-t.C:
			// ticker
			fmt.Println("ticker 5s")
		case i := <-ch:
			// channel
			fmt.Println("out", i)
		}
	}
}
