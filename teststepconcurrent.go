package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var v int = 99
var wg sync.WaitGroup
var s string

func Foo(x, y int) (z int) {
	//s = fmt.Sprintf("x=%d, y=%d, z=%d\n", x, y, z)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	z = x + y
	return
}

func Threads(fn func(x, y int) int) {
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 100; k++ {
				fn(1, 2)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			}
			wg.Done()
		}()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := v
	y := x * x
	var z int
	Threads(Foo)
	for i := 0; i < 100; i++ {
		z = Foo(x, y)
	}
	fmt.Printf("z=%d\n", z)
	wg.Wait()
}
