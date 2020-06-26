package goassessment

/**
 * Write a function, delayedCount, that takes two arguments:
 * a send channel 'c' and a receive channel 'q'.
 * The function is called within a goroutine.
 * The function should send the values 0 .. 4  on the send channel,
 * with a delay of 10 milliseconds or greater before sending each value.
 * The function should return when a message is received on the 'q' channel
 */
func delayedCount(c chan int, q chan int) {

}
