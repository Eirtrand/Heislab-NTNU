// Go 1.2
// go run helloworld_go.go
package main
import (
."fmt" 
"runtime"
"time"
)

var i int = 0

func someGoroutine1() {
	for k:=0;k<1000000;k++{
		i++
	}
}

func someGoroutine2() {
	for p:=0;p<1000000;p++{
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go someGoroutine1()
	go someGoroutine2() 
	time.Sleep(100*time.Millisecond)
	Println(i)
}


