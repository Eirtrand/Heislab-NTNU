package elevator

import (
	//"fmt"

	"time"
)

func Fsm() {
	InitializeElevator()

	for{
		time.Sleep(time.Millisecond * 18)
		PrintElev()
	}

}
