package main

//	func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
//	   ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
//	   defer cancel()
//	   p := processor{
//	       outA: make(chan AOut, 1),
//	       outB: make(chan BOut, 1),
//	       inC:  make(chan CIn, 1),
//	       outC: make(chan COut, 1),
//	       errs: make(chan error, 2),
//	   }
//	   p.launch(ctx, data)
//	   inputC, err := p.waitForAB(ctx)
//	   if err != nil {
//	       return COut{}, err
//	   }
//	   p.inC <- inputC
//	   out, err := p.waitForC(ctx)
//	   return out, err
//	}

func main() {

}
