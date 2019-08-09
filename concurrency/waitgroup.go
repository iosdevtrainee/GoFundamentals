package concurrency

import (
	"fmt"
	"sync"
)

func waitForGoroutines(num int) {
	if num <= 0 {
		return
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(num)
	for i := 0; i < num; i++ {
		go func(aNum int) {
			defer waitGroup.Done()
			fmt.Printf("goroutine number %d \n", aNum)
		}(i)
	}
	waitGroup.Wait()
}
