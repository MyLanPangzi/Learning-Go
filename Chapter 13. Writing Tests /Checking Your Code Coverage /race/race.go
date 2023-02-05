package race

import "sync"

func getCounter() int {
	var count int
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				count++
			}
		}()
	}
	wg.Wait()
	return count
}
