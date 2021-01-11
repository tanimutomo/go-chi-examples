package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSleep(ipt Input) Output {
	// Processing will take 1-5 seconds.
	rand.Seed(time.Now().Unix())
	processTime := time.Duration(rand.Intn(4)+1) * time.Second
	time.Sleep(processTime)

	return Output{
		Greeting: fmt.Sprintf("Hi, %s", ipt.Name),
	}
}

type Input struct {
	Name string
}

type Output struct {
	Greeting string
}
