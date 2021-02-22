package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tanimutomo/go-chi-examples/repository"
)

func GetGreeting(ctx context.Context, ipt Input) (Output, error) {
	fmt.Println(prefix, "start")
	begin := time.Now()

	greeting, err := repository.GetGreetingWord(ctx)
	if err != nil {
		return Output{}, err
	}

	end := time.Now()
	fmt.Println(prefix, "end")
	fmt.Println(prefix, "duration: ", end.Sub(begin).Seconds())

	return Output{
		Greeting: fmt.Sprintf("%s, %s", greeting, ipt.Name),
	}, nil
}
