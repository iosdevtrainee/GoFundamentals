package concurrency

import "fmt"

func producePositiveSequence(ofNums int, receiver chan<- int) {
	if ofNums < 1 {
		return
	}
	for i := 1; true; i++ {
		fmt.Printf("genNums num: %d \n", i)
		receiver <- i
	}
	close(receiver)
}

func pipelinePrintSquares(sender <-chan int) {
	for val := range sender {
		fmt.Printf("pipeline num: %d \n", val*val)
	}
}
