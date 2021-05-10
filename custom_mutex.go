package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type myMutex struct {
	isLocked int64 // 0 for unlocked, 1 for locked
}

func (m *myMutex) Lock() {
	// if it's not locked, change it to locked
	// if it's locked, keep looping until it's unlocked
	for atomic.CompareAndSwapInt64(&m.isLocked, 0, 1) == false {
		// do nothing and try again!
	}
}

func (m *myMutex) Unlock() {
	if atomic.CompareAndSwapInt64(&m.isLocked, 1, 0) == false {
		// error, mutex can't be unlocked because it isn't locked...
		// would panic in production
		fmt.Println("trying to unlock something we shouldn't...")
	}
}

type customCounter struct {
	count uint64
	mutex *myMutex
	cond *sync.Cond
}

var myCounter = customCounter{
	count: 0,
	mutex: &myMutex{},
	cond: sync.NewCond(&sync.Mutex{}),
}

func fakeWork() {
	sum := 0
	for i := 0; i < 50000; i++ {
		// do expensive stuff here
		sum += i
		if i % 2 == 0 {
			sum /= 2
			sum *= 7
		}
		for j := 1; j < 10000; j += 3 {
			sum = sum + j
			if sum > 1<<28 {
				sum /= 2
			}
		}
	}

}

func testCount(c *customCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i ++ {
		// we use cond when checking if the goroutine should "sleep" until unlock
		c.cond.L.Lock()
		// didn't see just an atomic compare, so swapping it with itself
		for atomic.CompareAndSwapInt64(&c.mutex.isLocked, 1, 1) {
			c.cond.Wait() // sleep while the mutex is held by a different go routine
		}

		c.mutex.Lock()
		c.count += 1
		// I empirically found sync.Cond to have about a 25% speedup with the presence of fakeWork
		// without fakeWork, sync.Cond slows down the program
		fakeWork()
		c.mutex.Unlock()

		// signal the lock is free to sleeping go routines
		c.cond.L.Unlock()
		c.cond.Signal()
	}
}

func main() {
	s := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go testCount(&myCounter, &wg)
	}
	wg.Wait()
	e := time.Now()
	fmt.Printf("time elapsed is: %d\n", e.Sub(s)) // prints time elapsed
	fmt.Printf("count is: %d\n", myCounter.count) // should print 100
}
