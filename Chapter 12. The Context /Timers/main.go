package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	now := time.Now()
	<-child.Done()
	fmt.Println(time.Since(now))
}
