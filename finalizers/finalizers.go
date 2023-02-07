package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

type Foo struct {
	a int
}

func main() {
	debug.SetGCPercent(-1)
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	for i := 0; i < 10; i++ {
		f := NewFoo(i)
		_ = fmt.Sprintf("%d", f.a)
	}

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		f := NewFoo(i)
		println(f.a)
	}
	runtime.GC()
	time.Sleep(1 * time.Second)
}

//go:noinline
func NewFoo(i int) *Foo {
	f := &Foo{a: rand.Intn(50)}
	runtime.SetFinalizer(f, func(f *Foo) {
		fmt.Println(`foo ` + strconv.Itoa(i) + ` has been garbage collected!`)
	})

	return f
}
