package driver

//import "fmt"

const n_BUTTONS int = 3
const n_FLOORS int = 4

 var lamp_channel_matrix = [n_FLOORS][n_BUTTONS] int{
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}


 var buttON_channel_matrix = [n_FLOORS][n_BUTTONS] int{
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func initialize_elevator() int {
	if !Init() {
		return 0
	}
	return 1
}

func poll_all_signals() {
	for {
		Get_floor_sensor_signal()
		get_button_signals()

		
	}
	
}

func Get_floor_sensor_signal() int{
	if Read_bit(SENSOR_FLOOR1) {
		return 1
	} else if Read_bit(SENSOR_FLOOR2) {
		return 2
	} else if Read_bit(SENSOR_FLOOR3) {
		return 3
	} else if Read_bit(SENSOR_FLOOR4) {
		return 4
	} else {
		return 0
	}
}

func get_button_signals() {
	
}

/*func power_off_all_lamps(){ 
	for i := 0; i<n_FLOORS; i++	{
		if i != 0 {
			elev_set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}

		if i != n_FLOORS - 1 {
			elev_set_button_lamp(BUTTON_CALL_UP, i, 0)
		}

		elev_set_button_lamp(BUTTON_COMMAND, i, 0)
	}
} */
