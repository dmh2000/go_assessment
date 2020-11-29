package goassessment

/**
 * Write a function, delayedCount, that takes two arguments:
 * a send channel 'c' and a receive channel 'q'.
 * The function is called within a goroutine.
 * The function should send the values 0 .. 4  on the send channel,
 * with a delay of 10 milliseconds or greater before sending each value.
 * When the function receives a message on the 'q' channel it should
 * send a final value of 10 then return
 */
func delayedCount(c chan int, q chan int) {
}
