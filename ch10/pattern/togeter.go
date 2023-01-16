package main

import "context"

type AOut struct {
}

type BOut struct {
}

type CIn struct {
}

type COut struct {
}

type processor struct {
	outA chan AOut
	outB chan BOut
	cIn  chan CIn
	outC chan COut
	errs chan error
}

type Input struct {
	B any
	A any
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()
	go func() {
		select {
		case <-ctx.Done():
			return
		case cIn := <-p.cIn:
			cOut, err := getResultC(ctx, cIn)
			if err != nil {
				p.errs <- err
			}
			p.outC <- cOut
		}
	}()
}

func getResultC(ctx context.Context, in CIn) (COut, error) {
	return COut{}, nil
}

func getResultB(ctx context.Context, b any) (BOut, error) {
	return BOut{}, nil
}

func getResultA(ctx context.Context, a any) (AOut, error) {
	return AOut{}, nil
}

func main() {
}
