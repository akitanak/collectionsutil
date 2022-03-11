package concurrent

import "sync"

func Map[IN, OUT any](l []IN, fn func(v IN) OUT) []OUT {
	mapped := make([]OUT, len(l))
	wg := sync.WaitGroup{}
	for i, v := range l {
		i, v := i, v
		wg.Add(1)
		go func() {
			defer wg.Done()
			mapped[i] = fn(v)
		}()
	}
	wg.Wait()
	return mapped
}
