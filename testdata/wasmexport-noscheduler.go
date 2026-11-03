package main

import (
	"io"
	"os"
	"time"
)

//go:noinline
func writeTo(w io.Writer, b []byte) {
	_, _ = w.Write(b)
}
func init() {
	println("called init")
}

//go:wasmimport tester callTestMain
func callTestMain()

func main() {
	// Check that exported functions can still be called after calling
	// time.Sleep.
	time.Sleep(time.Millisecond)

	// main.main is not used when using -buildmode=c-shared.
	callTestMain()
}

//go:wasmexport hello
func hello() {
	writeTo(io.Discard, []byte("discard"))
	writeTo(os.Stdout, nil)
	println("hello!")
}

//go:wasmexport add
func add(a, b int32) int32 {
	println("called add:", a, b)
	return a + b
}

//go:wasmimport tester callOutside
func callOutside(a, b int32) int32

//go:wasmexport reentrantCall
func reentrantCall(a, b int32) int32 {
	println("reentrantCall:", a, b)
	result := callOutside(a, b)
	println("reentrantCall result:", result)
	return result
}

//go:wasmexport goroutineExit
func goroutineExit() {
	// Dummy, not a real test (since we have no scheduler).
	println("goroutineExit: exit")
}
