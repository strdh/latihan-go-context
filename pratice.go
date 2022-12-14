package main

import (
	"fmt"
	"time"
	"context"
	"runtime"
)

func main() {
	sum := make(chan int)
	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent)
	// ctx, cancel := context.WithTimeout(parent, 6*time.Second)
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(6*time.Second))
	defer cancel()

	go func() {
		defer close(sum)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancel dijalankan")
				return
			default:
				sum <- counter
				counter++
				time.Sleep(2 * time.Second)
			}
		}
	}()

	fmt.Println("Jumlah Go-Routine : ", runtime.NumGoroutine())

	//with cancel
	// for iter := range sum {
	// 	fmt.Println("Counter : ", iter)
	// 	if iter == 10 {
	// 		cancel()
	// 		// break
	// 	}
	// }

	//with timeout
	for iter := range sum {
		fmt.Println("Counter : ", iter)
	}

	fmt.Println("Jumlah Go-Routine : ", runtime.NumGoroutine())


}