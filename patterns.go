package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concurrency patterns")

	// channel := generateChannel("hello")

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-channel)
	// }

	// Fan in pattern
	//fanInChannel := fanIn(generateChannel("hi"), generateChannel("hello"))

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-fanInChannel)
	// }

	// worker pools
	//compute([]int{1, 2, 3, 4, 5, 6}, 3, 10)

	// fan-in / fan-out
	//computeResult()

	//doSemaphoreExecution()

	doTimeoutExecution()

	doTImeoutWithCOntext()
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

// Semaphore pattern

func doSemaphoreExecution() {
	var wg sync.WaitGroup

	sem := make(chan struct{}, 5)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go accessDatabase(i+1, sem, &wg)
	}

	wg.Wait()
	fmt.Println("All 10 goroutines completed their task")
}

// A function simulating database access
func accessDatabase(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{}
	fmt.Printf("Goroutine %d: Accessing database\n", id)

	time.Sleep(time.Second * 2)

	<-sem
	fmt.Printf("Goroutine %d: Completed with database access\n", id)
}

// Timeout pattern

func fetchData() string {
	time.Sleep(5 * time.Second)
	return "Data"
}

func fetchDataWithTimeout(timeout time.Duration) (string, error) {
	result := make(chan string)
	go func() {
		result <- fetchData()
	}()

	select {
	case res := <-result:
		return res, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("operation exited with timeout")
	}
}

func doTimeoutExecution() {
	timeout := 10 * time.Second

	res, err := fetchDataWithTimeout(timeout)

	if err != nil {
		fmt.Println("Error timeout : ", err)
	} else {
		fmt.Println("Result timeout : ", res)
	}
}

func fetchDataWithContext(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		return "Data from server", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchWithTimeout(timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	return fetchDataWithContext(ctx)
}

func doTImeoutWithCOntext() {
	timeout := 6 * time.Second

	res, err := fetchWithTimeout(timeout)

	if err != nil {
		fmt.Println("Error ctx : ", err)
	} else {
		fmt.Println("Res ctx : ", res)
	}
}
