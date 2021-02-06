package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

const prefix = "    [Usecase] "

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
