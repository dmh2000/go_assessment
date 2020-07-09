package goassessment

import (
	"sync"
	"testing"
	"time"
)

/**
 * Write a function, delayedCount, that takes two arguments:
 * a send channel 'c' and a receive channel 'q'.
 * The function is called within a goroutine.
 * The function should send the values 0 .. 4  on the send channel,
 * with a delay of 10 milliseconds or greater before sending each value.
 * The function should return when a message is received on the 'q' channel
 */
func TestAsync(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors

	t.Log("you should be able to send values to a channel with a delay")

	var c chan int
	var q chan int
	var wg sync.WaitGroup

	c = make(chan int)
	q = make(chan int, 2)

	wg.Add(2)

	go func() {
		var j int
	loop:
		for i := 0; i < 5; i++ {
			t0 := time.Now()

			select {
			case j = <-c:
				if j != i {
					t.Error(shouldBe(j, i))
				}
			case <-time.After(100 * time.Microsecond):
				t.Error("select statement timed out")
				break loop
			}

			// compute time difference
			t1 := time.Now()
			d := t1.Sub(t0)
			if d.Milliseconds() < 10 {
				t.Error("delay of ", d.Milliseconds(), "should be greater than 10ms")
			}
		}
		q <- 0
		wg.Done()
	}()

	go func() {
		delayedCount(c, q)
		wg.Done()
	}()

	wg.Wait()
}
