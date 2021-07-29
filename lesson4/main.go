package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var workers = make(chan int, 10)
	var a = 0
	var m sync.Mutex

	for i := 1; i <= 1000; i++ {
		workers <- i

		go func(job int) {
			defer func() {
				<-workers
			}()
			m.Lock()
			a++
			m.Unlock()

		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
