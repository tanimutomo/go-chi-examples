package repository

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const prefix = "      [Repository] "

func DoSomethingWithRandomSleep(ctx context.Context) error {
	done := make(chan int)
	go func() {
		doSomethingWithRandomSleep(ctx)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println(prefix, "timeout")
		return errors.New("Timeout")
	case <-done:
		return nil
	}
}

func doSomethingWithRandomSleep(ctx context.Context) {
	defer func() {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println(prefix, "rollback")
		} else {
			fmt.Println(prefix, "commit")
		}
	}()

	fmt.Println(prefix, "start")
	begin := time.Now()

	// Processing will take 1-5 seconds.
	rand.Seed(time.Now().Unix())
	waitSec := rand.Intn(4) + 1
	processTime := time.Duration(waitSec) * time.Second
	time.Sleep(processTime)

	end := time.Now()
	fmt.Println(prefix, "end")
	fmt.Println(prefix, "duration: ", end.Sub(begin).Seconds())
}
