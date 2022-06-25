package concurrent

import "sync"

// Map maps input slice values to output slice values `l` with specified function `fn` concurrently.
func Map[IN, OUT any](s []IN, fn func(v IN) OUT) []OUT {
	mapped := make([]OUT, len(s))
	wg := sync.WaitGroup{}
	for i, v := range s {
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

// FlatMap maps input slice values `l` with specified function `fn` that returns slice and flatten slices concurrently.
func FlatMap[IN, OUT any](s []IN, fn func(v IN) []OUT) []OUT {
	mutex := &sync.Mutex{}
	mapped := make(map[int][]OUT, 0)
	wg := sync.WaitGroup{}
	for i, v := range s {
		i, v := i, v
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			mapped[i] = fn(v)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	flatten := make([]OUT, 0)
	for i := 0; i < len(s); i++ {
		flatten = append(flatten, mapped[i]...)
	}
	return flatten
}
