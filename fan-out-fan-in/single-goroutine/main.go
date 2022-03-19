package main

import (
	"fmt"
	"sync"
	"time"
)

//The fan-out-fan-in pattern is a modification on the pipeline pattern so handle slow load processing.
//There could be scenarios where one of the stage in the pipeline consumes a lot of time, it would be better if this stage was run via multiple goroutines instead of one.
//This way , the load is processed faster.
//This in essence is the fan-out pattern , where multiple goroutines read from a common channel and write out to multiple channels thus "faning out".
//The fan-in would now read from multiple channels into a single channel and process the load.

//The task of computing squares has a sleep of 5 seconds added to simulate work.
//This example just generates 1 goroutine for processing the worker task(i.e the square computing task)
//So even if this is a fan-in-fan-out design , it is basically a pipeline and it will take the same time each task takes to finish up.
//In this case , we are generating square of 10 numbers , each computation sleeps for a duration of 5 seconds , so total time taken will be 50.
//Run the program and observe the output.
func main() {
	start := time.Now()
	gnc := generateNumbers()
	csc1 := computeSquares(gnc)

	fanInChannel := merge([]<-chan int{csc1})

	for fic := range fanInChannel {
		fmt.Println("Value from fan in channel", fic)
	}

	fmt.Println("DOne in ", time.Since(start))
}

func generateNumbers() <-chan int {
	generateNumberChannel := make(chan int)
	go func() {
		defer close(generateNumberChannel)
		for i := 0; i < 10; i++ {
			generateNumberChannel <- i
		}
	}()
	return generateNumberChannel
}

func computeSquares(generateNumberChannel <-chan int) <-chan int {
	computeSquaresChannel := make(chan int)
	go func() {
		defer close(computeSquaresChannel)
		for ch := range generateNumberChannel {
			time.Sleep(time.Second * 5)
			computeSquaresChannel <- ch * ch
		}
	}()
	return computeSquaresChannel
}

//gets multiple read only channels
//merges the channels and writes the output to a single channel which can then be used for processing.
func merge(channels []<-chan int ) <-chan int {
	fanInChannel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			for c1 := range c {
				fanInChannel <- c1
			}
			wg.Done()
		}(c)
	}
	go func(){
		wg.Wait()
		defer close(fanInChannel)
	}()
	return fanInChannel
}
