package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	jobs := make(chan int64, 50)
	result := make(chan int64, 50)

	go worker(1, jobs, result)
	go worker(2, jobs, result)
	go worker(3, jobs, result)

	for i := 0; i < 50; i++ {
		jobs <- int64(i)
	}
	close(jobs)

	for i := 0; i < 50; i++ {
		fmt.Println(<-result)
	}
}

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(id int64, jobChan <-chan int64, resultChan chan<- int64) {
	for job := range jobChan {
		startTime := time.Now()
		log.Printf(
			"[%s] Worker with ID:%d is starting to execute job n=%d\n",
			startTime.Format("2006-01-02T15:04:05.999999-07:00"),
			id,
			job,
		)
		res := fibonacci(job)
		resultChan <- res
		endTime := time.Now()
		log.Printf(
			"[%s] Worker with ID:%d has finished executing job n=%d, result=%d\n",
			endTime.Format("2006-01-02T15:04:05.999999-07:00"),
			id,
			job,
			res,
		)

	}

}
