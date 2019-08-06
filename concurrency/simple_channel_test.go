package concurrency

import (
	"testing"
)

func TestProducerChannel(t *testing.T) {
	channel := make(chan int)
	// channels block when writing to a channel
	// requires a go in front of consumer channels
	go producePositiveSequence(4, channel)
	// var values []int
	// intSet := []int{1, 2, 3, 4}
	// for val := range channel {
	// 	values = append(values, val)
	// }
	pipelinePrintSquares(channel)
	// for index, val := range intSet {
	// 	if val != values[index] {
	// 		t.Fail()
	// 	}
	// }

}
