package main

import (
	"sync/atomic"
	"testing"
	"sync"
)

type Race struct {
	Value string
}

func BenchmarkRaceCondition(b *testing.B) {
	mu := sync.Mutex{}
	race := &Race{
		Value: "init",
	}
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			race = &Race{
				Value: "first",
			}
			mu.Unlock()
		}()
		mu.Lock()
		race = &Race{
			Value: "second",
		}
		mu.Unlock()
	}
	b.StopTimer()
	b.Log("race value", race.Value)
}

func BenchmarkConditionAtomicPointer(b *testing.B) {
	race := atomic.Pointer[Race]{}
	race.Store(&Race{
		Value: "init",
	})
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		go func() {
			race.Swap(&Race{
				Value: "first",
			})
		}()
		{
			race.Swap(&Race{
				Value: "second",
			})
		}
	}
	b.StopTimer()

	b.Log("race value", *race.Load())
}