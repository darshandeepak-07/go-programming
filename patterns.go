package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concurrency patterns")

	channel := generateChannel("hello")

	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}

	// Fan in pattern
	fanInChannel := fanIn(generateChannel("hi"), generateChannel("hello"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-fanInChannel)
	}

	// worker pools
	//compute([]int{1, 2, 3, 4, 5, 6}, 3, 10)

	// fan-in / fan-out
	computeResult()
}

// Generator Pattern - returns a channel

func generateChannel(msg string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("%s : %d\n", msg, i)
		}
	}()
	return channel
}

// Fan in

func fanIn(channel1, channel2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-channel1
		}
	}()

	go func() {
		for {
			c <- <-channel2
		}
	}()
	return c
}

// Function to implement worker pool
// Calculating square of numbers using worker pool pattern

func compute(list []int, noOfWorker, noOfJobs int) {
	jobs := make(chan int, noOfJobs)
	output := make(chan int, noOfJobs)
	var wg sync.WaitGroup

	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go square(i+1, jobs, output, &wg)
	}

	for _, num := range list {
		jobs <- num
	}
	close(jobs)

	wg.Wait()
	close(output)

	for res := range output {
		fmt.Println(res)
	}
}

// worker function to calculate square
func square(workedId int, jobs <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range jobs {
		square := num * num
		fmt.Printf("Worker %d : Result = %d\n", workedId, square)
		time.Sleep(time.Second)
		output <- square
	}
}

// Fan-In and Fan-out Implemetation
// Generating random numbers process them in multiple workers and merge their result in
// single output channel

func computeResult() {
	noJobs := 10
	noWorkers := 3

	jobs := generateRandomNumbers(noJobs)
	result := make([]chan int, noWorkers)
	var wg sync.WaitGroup

	for i := 0; i < noWorkers; i++ {
		result[i] = make(chan int, noJobs)
		wg.Add(1)
		go square(i+1, jobs, result[i], &wg)
	}

	go func() {
		wg.Wait()
		for _, ch := range result {
			close(ch)
		}
	}()

	mergedResult := fanInChannel(result...)

	for res1 := range mergedResult {
		fmt.Println("Res = ", res1)
	}
}

func generateRandomNumbers(count int) <-chan int {
	out := make(chan int, count)

	go func() {
		for i := 0; i < count; i++ {
			num := rand.Intn(11)
			fmt.Println("Generated num : ", num)
			out <- num
		}
		close(out)
	}()

	return out
}

func fanInChannel(channels ...chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, channel := range channels {
		wg.Add(1)

		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
