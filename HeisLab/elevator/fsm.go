package elevator

import (
	"fmt"
	"../driver"
	"time"
)

func Fsm() {
	for{
		 fmt.Println(driver.Get_floor_sensor_signal())
		 time.Sleep(time.Millisecond * 100)
	}

}
