// Go 1.2
// go run helloworld_go.go
package main

import (
	. "fmt"
	"runtime"
	"time"
)

var done_add = make(chan string)
var cs = make(chan int, 1)
var i int = 0

func someGoroutine1() {
	for k := 0; k < 1000000; k++ {
		<-cs
		i++
		cs <- 1
	}
	done_add <- "done adding"
}

func someGoroutine2() {
	for p := 0; p < 1000001; p++ {
		<-cs
		i--
		cs <- 1
	}
	done_add <- "done subtracting"
}

func main() {
	cs <- 1
	runtime.GOMAXPROCS(runtime.NumCPU())
	go someGoroutine1()
	go someGoroutine2()
	time.Sleep(100 * time.Millisecond)
	Println(<-done_add)
	Println(<-done_add)
	Println(i)
}
