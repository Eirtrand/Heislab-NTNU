package elevator

import (
	"fmt"
	"../driver"
	"time"
)

func Fsm() {
	for{
		 //fmt.Println(driver.Get_floor_sensor_signal())
		fmt.Println("Light ON")
		driver.Elev_set_door_open_lamp(1)
		 time.Sleep(time.Millisecond * 2000)

		 fmt.Println("Light OFF")
		 driver.Elev_set_door_open_lamp(0)
		 time.Sleep(time.Millisecond * 2000)
	}

}
