package main

import (
	"fmt"
	"time"
)

// Fibonacci function
func Fibonacci(n int, cache *Memory) int {
	if n <= 1 {
		return n
	}
	f1, _ := cache.Get(n - 1)
	f2, _ := cache.Get(n - 2)
	return f1.(int) + f2.(int)
}

// cache
type Memory struct {
	f     Function
	cache map[int]FuncResult
}
type Function func(key int, cache *Memory) (interface{}, error)

type FuncResult struct {
	value interface{}
	err   error
}

// NewCache function
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FuncResult),
	}
}

// Get function
func (m *Memory) Get(key int) (interface{}, error) {
	if res, ok := m.cache[key]; ok {
		return res.value, res.err // return from cache
	}
	res := FuncResult{}
	res.value, res.err = m.f(key, m) // call the function
	m.cache[key] = res               // store the result in cache
	return res.value, res.err
}

// GetFibonacci function
func GetFibonacci(n int, cache *Memory) (interface{}, error) {
	return Fibonacci(n, cache), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38, 40, 40, 41, 43, 43, 43}
	for _, v := range fibo {
		start := time.Now()
		result, err := cache.Get(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d,Time taken: %v, result: %d \n", v, time.Since(start), result)
	}
}
