package main

import (
	"fmt"
	"time"
)

//In a pipeline pattern , output of the previous stage serves as input to the next.
//The code is a simple pipeline pattern to which consists of 3 stages.
// Generate a number ---> Compute square of the number ---> Print the number.
//This simulates a producer , some consumer and some final output pattern.
//A sleep is added to simulate work done and more importantly to see what difference a buffered and a non-buffered channel makes.
//Observe the print statements added. Every write action waits for the sender to read the from the channel.
//It simulates a scenario where sender is fast but consumer is slow and having non-buffered channel for the sender, might not be a good idea.
//This is a non-buffered channel , so the sender will have to wait till a receiver has picked up the written value from the queue.
//Example of buffered channel can be found in a separate file.
func main() {

	start := time.Now()
	for c:= range computeSquares(generateNumbers()){
		fmt.Println("Square root is " ,c)
	}
	fmt.Println("total time elapsed", time.Since(start))
}

func generateNumbers() <-chan int {
	c := make(chan int)
	go func() {
		defer func() {
			close(c)
		}()
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("Wrote ", i ," ..waiting for another worker thread to read from this channel")
		}
	}()
	return c
}

func computeSquares(generateNumberChannel <-chan int) <-chan int {
	computeChannel := make(chan int)
	go func() {
		defer func() {
			close(computeChannel)
		}()
		for gnc := range generateNumberChannel {
			computeChannel <- gnc*gnc
			time.Sleep(time.Second *3)
		}
	}()
	return computeChannel
}