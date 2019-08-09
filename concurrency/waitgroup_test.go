package concurrency

import "testing"

func TestWaitGroup(t *testing.T) {
	waitForGoroutines(10)
}
