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
 * When the function receives a message on the 'q' channel it should
 * send a final value of 10 then return
 */

func TestAsync(t *testing.T) {
	defer testPanic(t) // handle panics and syntax errors // handle panics and syntax errors
	setTestInstance(t) // update test instance for logging

	t.Log("GOAL: you should be able to send values to a channel with a delay")

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
			// record current time
			t0 := time.Now()

			// wait for response from go routine or timeout
			select {
			case j = <-c:
				if j != i {
					t.Error(shouldBe(j, i))
				}
			case <-time.After(2000 * time.Millisecond):
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
		// signal quit
		q <- 0

		// check final value
		j = <-c
		if j != 10 {
			t.Error(shouldBe(j, 10))
		}
		wg.Done()
	}()

	go func() {
		delayedCount(c, q)
		wg.Done()
	}()

	wg.Wait()
}
