package main

import (
	"fmt"
	"sync"
	"testing"
)

type Set struct {
	a int
	sync.RWMutex
}

func (s *Set) Add() {
	defer s.Unlock()
	s.Lock()
	s.a++
}

func (s *Set) Has() int {
	defer s.RUnlock()
	s.RLock()
	return s.a
}

func worker(s *Set, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	if id <= 9000 {
		s.Add()
	}
	fmt.Printf("Worker %d, a: %d\n", id, s.Has())
}

func main() {
	bs := testing.Benchmark(BenchmarkStart)
	br := testing.Benchmark(BenchmarkAdd)
	brh := testing.Benchmark(BenchmarkHas)
	fmt.Println(bs)
	fmt.Println(br)
	fmt.Println(brh)
}

func start() {
	var wg sync.WaitGroup
	var s = &Set{}
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go worker(s, i, &wg)
	}
	wg.Wait()
	fmt.Println(s.Has())
}

func BenchmarkAdd(b *testing.B) {
	var set = &Set{}
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add()
			}
		})
	})

}

func BenchmarkHas(b *testing.B) {
	var set = &Set{}
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has()
			}
		})
	})
}

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start()
	}
}
