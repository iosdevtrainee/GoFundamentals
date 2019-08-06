package concurrency

import (
	"fmt"
	"runtime"
)

func printMeAsync(data string) {
	for range data {
		fmt.Println(data)
	}
}

// no guarantee the goroutine will run because
// the runAsync may complete before the goroutine has
// a chance to complete
func runAsync(myFunc func(string)) {
	go myFunc("data")
}

// goroutine has been scheduled to run because you've
// indicated to the scheduler that you have time to run
// other tasks in the queue
func runAsyncGuarantee(myFunc func(string)) {
	go myFunc("data")
	runtime.Gosched()
}
