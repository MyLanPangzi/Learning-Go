package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Param struct {
	a int
	b int
}

func main() {
	r := GatherAndProcess(context.Background(), Param{1, 1})
	if r.err != nil {
		log.Fatal(r.err)
	}
	fmt.Println(r)
}

type ResultC struct {
	r   int
	err error
}

type ParamC struct {
	a int
	b int
}

func GatherAndProcess(ctx context.Context, param Param) ResultC {
	//ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1000)
	defer cancel()
	aStream := multiplyTwo(param.a)
	bStream := addOne(param.b)
	var paramC ParamC
	for i := 0; i < 2; {
		select {
		case a := <-aStream:
			paramC.a = a.r
			aStream = nil
			i++
		case b := <-bStream:
			paramC.b = b.r
			bStream = nil
			i++
		case <-ctx.Done():
			return ResultC{err: ctx.Err()}
		}
	}
	cStream := processC(paramC)
	select {
	case out := <-cStream:
		return out
	case <-ctx.Done():
		return ResultC{err: ctx.Err()}
	}
}

func processC(in ParamC) <-chan ResultC {
	stream := make(chan ResultC)
	go func() {
		defer close(stream)
		time.Sleep(time.Second)
		stream <- ResultC{
			r:   in.a + in.b,
			err: nil,
		}
	}()
	return stream
}

type ResultB struct {
	r   int
	err error
}

func addOne(b int) <-chan ResultB {
	stream := make(chan ResultB)
	go func() {
		defer close(stream)
		time.Sleep(time.Second)
		stream <- ResultB{
			r:   b + 1,
			err: nil,
		}
	}()
	return stream
}

type ResultA struct {
	r   int
	err error
}

func multiplyTwo(a int) <-chan ResultA {
	stream := make(chan ResultA)
	go func() {
		defer close(stream)
		time.Sleep(time.Second)
		stream <- ResultA{
			r:   a * 2,
			err: nil,
		}
	}()
	return stream
}
