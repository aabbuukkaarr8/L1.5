package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := 3 * time.Second

	done := time.After(timeout)

	values := make(chan int)

	go startProducer(values, done, 200*time.Millisecond)

	startConsumer(values, done)
}

func startProducer(out chan<- int, done <-chan time.Time, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer close(out)

	value := 0
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			select {
			case out <- value:
				value++
			case <-done:
				return
			}
		}
	}
}

func startConsumer(in <-chan int, done <-chan time.Time) {
	for {
		select {
		case <-done:
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}
