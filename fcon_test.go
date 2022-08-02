package main

import (
	"fmt"
	"context"
	"testing"
	"runtime"
	"time"
)

// func TestContext(t *testing.T) {
// 	ctx := context.Background()
// 	fmt.Println(ctx)

// 	todo := context.TODO()
// 	fmt.Println(todo)
// }

// func TestContextWIthValue(t *testing.T) {
// 	ctx1 := context.Background()
// 	ctxA := context.WithValue(ctx1, "a", "A")
// 	ctxB := context.WithValue(ctx1, "b", "B")

// 	ctxC := context.WithValue(ctxA, "c", "C")
// 	ctxD := context.WithValue(ctxA, "d", "D")

// 	ctxE := context.WithValue(ctxB, "e", "E")

// 	fmt.Println(ctxA)
// 	fmt.Println(ctxB)
// 	fmt.Println(ctxC)
// 	fmt.Println(ctxD)
// 	fmt.Println(ctxE.Value("b"))
// }

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Off")
				return
			default:
				destination <-counter 
				counter++
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return destination
}

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	parent := context.Background()
// 	ctx, cancel := context.WithCancel(parent)

// 	destination := CreateCounter(ctx)
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// 	for n := range destination {
// 		fmt.Println("Counter : ", n)
// 		if n == 10 {
// 			cancel()
// 			// break
// 		}
// 	}
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

// func TestContextWithTimeout(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	parent := context.Background()
// 	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
// 	defer cancel()

// 	destination := CreateCounter(ctx)
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// 	for n := range destination {
// 		fmt.Println("Counter : ", n)
// 	}
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter : ", n)
	}
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}