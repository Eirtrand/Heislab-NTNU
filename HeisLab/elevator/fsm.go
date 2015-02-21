package elevator

import (
	"fmt"
	"../driver"
	"time"
)

func Fsm() {
	driver.InitializeElevator()

	for{
		fmt.Println(driver.STOP == 790)
		time.Sleep(time.Millisecond * 18)

	}

}
