package driver

import "fmt"

func initialize_elevator(chan floor_chan int, chan button_chan int) {
	if !Init() int {
		return 0
	}
	for{
		select{
			case
		}
	}

}

func poll_all_signals() {
	for {
		get_floor_sensor()
		get_button_signals()

		
	}
	
}

func get_floor_sensor_signal() {
	if Read_bit(SENSOR_FLOOR1) {
		return 1
	}
	else if Read_bit(SENSOR_FLOOR2) {
		return 2
	}
	else if Read_bit(SENSOR_FLOOR3) {
		return 3
	}
	else if Read_bit(SENSOR_FLOOR4) {
		return 4
	}
	else {
		return 0
	}
	
}

func get_button_signals() {
	
}