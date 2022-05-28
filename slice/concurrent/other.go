package concurrent

import "sync"

// ForEach apply specified function to specified slice values concurrently.
func ForEach[T any](s []T, fn func(e T)) {
	wg := sync.WaitGroup{}
	for _, e := range s {
		e := e
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(e)
		}()
	}
	wg.Wait()
}
