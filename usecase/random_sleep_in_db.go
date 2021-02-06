package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tanimutomo/go-chi-examples/repository"
)

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
