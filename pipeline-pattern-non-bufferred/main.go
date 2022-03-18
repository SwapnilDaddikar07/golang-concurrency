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
//Observe the print statements added.

// The sender and receiver are now using a non-buffered channel. The sender will benefit from this because the receiver is slow. (time delay of 3 seconds)
//It simulates a scenario where sender is fast but consumer is slow and having non-buffered channel for the sender, might not be a good idea.
//This is a buffered channel example , so the sender will have dump all the data at once , it wont wait for the receiver to pull out data from channel.
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
			fmt.Println("Wrote ", i)
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