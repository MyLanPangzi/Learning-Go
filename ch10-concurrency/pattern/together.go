package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type processor struct {
	outA   chan int
	outB   chan int
	inC    chan InputC
	outC   chan int
	errors chan error
}

type Input struct {
	a int
	b int
}

func GatherAdnProcess(ctx context.Context, data Input) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()
	p := processor{
		outA:   make(chan int, 1),
		outB:   make(chan int, 1),
		inC:    make(chan InputC, 1),
		outC:   make(chan int, 1),
		errors: make(chan error, 2),
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return 0, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

func (p processor) launch(ctx context.Context, data Input) {
	go func() {
		out, err := getResultA(ctx, data.a)
		if err != nil {
			p.errors <- err
			return
		}
		p.outA <- out
	}()
	go func() {
		out, err := getResultB(ctx, data.b)
		if err != nil {
			p.errors <- err
			return
		}
		p.outB <- out
	}()
	go func() {
		select {
		case <-ctx.Done():
			return
		case in := <-p.inC:
			out, err := getResultC(ctx, in)
			if err != nil {
				p.errors <- err
				return
			}
			p.outC <- out
		}
	}()
}

type InputC struct {
	a int
	b int
}

func getResultC(ctx context.Context, in InputC) (int, error) {
	time.Sleep(time.Second)
	return in.b + in.a, nil
}

func getResultB(ctx context.Context, b int) (int, error) {
	time.Sleep(time.Second)
	return 2 * b, nil
}

func getResultA(ctx context.Context, a int) (int, error) {
	time.Sleep(time.Second)
	return 1 + a, nil
}

func (p processor) waitForAB(ctx context.Context) (InputC, error) {
	var out InputC
	for i := 0; i < 2; {
		select {
		case a := <-p.outA:
			i++
			out.a = a
		case b := <-p.outB:
			i++
			out.b = b
		case err := <-p.errors:
			return InputC{}, err
		case <-ctx.Done():
			return InputC{}, ctx.Err()
		}
	}
	return out, nil
}

func (p processor) waitForC(ctx context.Context) (int, error) {
	select {
	case c := <-p.outC:
		return c, nil
	case err := <-p.errors:
		return 0, err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
func main() {
	r, err := GatherAdnProcess(context.Background(), Input{
		a: 1,
		b: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
