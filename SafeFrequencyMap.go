package main

import "sync"

type SafeFrequencyMap struct {
	mu sync.Mutex
	mp map[string]int // Actual map
}

func (sfmp *SafeFrequencyMap) IncrementFrequency(key string) {
	sfmp.mu.Lock() // Lock acquire attempt (Blocking)
	defer sfmp.mu.Unlock()

	_, ok := sfmp.mp[key]
	if !ok {
		sfmp.mp[key] = 1
	} else {
		sfmp.mp[key]++
	}
}
