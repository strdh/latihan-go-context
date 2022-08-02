package main

import (
	"fmt"
	// "time"
	"context"
	"runtime"
)

func main() {
	sum := make(chan int)
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

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
			}
		}
	}()

	fmt.Println("Jumlah Go-Routine : ", runtime.NumGoroutine())

	for iter := range sum {
		fmt.Println("Counter : ", iter)
		if iter == 10 {
			cancel()
			// break
		}
	}

	fmt.Println("Jumlah Go-Routine : ", runtime.NumGoroutine())


}