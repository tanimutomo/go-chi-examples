package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/tanimutomo/go-chi-examples/repository"
)

const prefix = "    [Usecase] "

type Input struct {
	Name string
}

type Output struct {
	Greeting string
}

func RandomSleep(ipt Input) Output {
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

	return Output{
		Greeting: fmt.Sprintf("Hi, %s", ipt.Name),
	}
}

func RandomSleepInDB(ctx context.Context, ipt Input) (Output, error) {
	fmt.Println(prefix, "start")
	begin := time.Now()

	err := repository.DoSomethingWithRandomSleep(ctx)
	if err != nil {
		return Output{}, err
	}

	end := time.Now()
	fmt.Println(prefix, "end")
	fmt.Println(prefix, "duration: ", end.Sub(begin).Seconds())

	return Output{
		Greeting: fmt.Sprintf("Hi, %s", ipt.Name),
	}, nil
}
