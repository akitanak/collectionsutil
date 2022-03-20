package concurrent

import "sync"

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
