package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSleep(ipt Input) Output {
	fmt.Println("    [Usecase] start")
	begin := time.Now()

	// Processing will take 1-5 seconds.
	rand.Seed(time.Now().Unix())
	waitSec := rand.Intn(4) + 1
	processTime := time.Duration(waitSec) * time.Second
	time.Sleep(processTime)

	end := time.Now()
	fmt.Println("    [Usecase] end")
	fmt.Println("    [Usecase] duration: ", end.Sub(begin).Seconds())

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
