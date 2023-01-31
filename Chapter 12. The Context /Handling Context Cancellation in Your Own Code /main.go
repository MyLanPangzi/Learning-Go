package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	data, err := longRunningThingManager(ctx, "world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		r   string
		err error
	}
	valStream := make(chan wrapper, 1)
	go func() {
		defer close(valStream)
		r, err := doSomeThing(ctx, data)
		valStream <- wrapper{r, err}
	}()
	select {
	case val := <-valStream:
		return val.r, val.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func doSomeThing(ctx context.Context, data string) (string, error) {
	time.Sleep(2 * time.Second)
	return "hello " + data, nil
}
