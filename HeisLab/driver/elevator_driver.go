package driver

import(
) 
const n_BUTTONS int = 3
const n_FLOORS int = 4

type State int
var elev elevator
type Direction int

const (
	NONE Direction = iota
	UP 
	DOWN 
)


const (
	INIT State = iota
	STOPPED
 	GOING_UP
 	GOING_DOWN
	EMERGENCY_STOP
)

 var external_button_array = [NUMBER_OF_EXT_BUTTONS] int{
 	BUTTON_UP1, BUTTON_UP2, BUTTON_UP3 ,BUTTON_DOWN2, BUTTON_DOWN3, BUTTON_DOWN4,
}

 var internal_button_array = [NUMBER_OF_INT_BUTTONS] int{
 	BUTTON_COMMAND1, BUTTON_COMMAND2, BUTTON_COMMAND3, BUTTON_COMMAND4,
}

type elevator struct {
	state      State
	externalButtons []int
	internalButtons []int
	stop bool
	obstruction bool
	currentFloor int
	prevFloor int
	prevDirection Direction
}


func GetFloorSensorSignal() int{
	if ReadBit(SENSOR_FLOOR1) {
		return 1
	} else if ReadBit(SENSOR_FLOOR2) {
		return 2
	} else if ReadBit(SENSOR_FLOOR3) {
		return 3
	} else if ReadBit(SENSOR_FLOOR4) {
		return 4
	} else {
		return -1
	}
}

func SetDoorOpenLamp(value int) {
	if value == 1{
		SetBit(LIGHT_DOOR_OPEN)
	}else{
		ClearBit(LIGHT_DOOR_OPEN)
	}
}

func SetLight(floor int, dir Direction) {
	switch {
	case floor == 1 && dir == NONE:
		SetBit(LIGHT_COMMAND1)
	case floor == 2 && dir == NONE:
		SetBit(LIGHT_COMMAND2)
	case floor == 3 && dir == NONE:
		SetBit(LIGHT_COMMAND3)
	case floor == 4 && dir == NONE:
		SetBit(LIGHT_COMMAND4)
	case floor == 1 && dir == UP:
		SetBit(LIGHT_UP1)
	case floor == 2 && dir == UP:
		SetBit(LIGHT_UP2)
	case floor == 3 && dir == UP:
		SetBit(LIGHT_UP3)
	case floor == 2 && dir == DOWN:
		SetBit(LIGHT_DOWN2)
	case floor == 3 && dir == DOWN:
		SetBit(LIGHT_DOWN3)
	case floor == 4 && dir == DOWN:
		SetBit(LIGHT_DOWN4)
	}
}

func ClearLight(floor int, dir Direction) {
	switch {
	case floor == 1 && dir == NONE:
		ClearBit(LIGHT_COMMAND1)
	case floor == 2 && dir == NONE:
		ClearBit(LIGHT_COMMAND2)
	case floor == 3 && dir == NONE:
		ClearBit(LIGHT_COMMAND3)
	case floor == 4 && dir == NONE:
		ClearBit(LIGHT_COMMAND4)
	case floor == 1 && dir == UP:
		ClearBit(LIGHT_UP1)
	case floor == 2 && dir == UP:
		ClearBit(LIGHT_UP2)
	case floor == 3 && dir == UP:
		ClearBit(LIGHT_UP3)
	case floor == 2 && dir == DOWN:
		ClearBit(LIGHT_DOWN2)
	case floor == 3 && dir == DOWN:
		ClearBit(LIGHT_DOWN3)
	case floor == 4 && dir == DOWN:
		ClearBit(LIGHT_DOWN4)
	}
}

func SetMotorDirection(dir Direction) {
	switch{
 	case dir == NONE:
		WriteAnalog(MOTOR, 0);
	case dir == UP: 
		ClearBit(MOTORDIR);
		WriteAnalog(MOTOR, 2800);
	case dir == DOWN:
		SetBit(MOTORDIR);
		WriteAnalog(MOTOR, 2800);
	}
}

func GetObstructionSignal() bool {
	return ReadBit(STOP)
}

func GetButtonSignal(buttonType int) bool {
	return ReadBit(buttonType)
}

/*

*/
func setFloorLight(floor int){
	if floor == 0x02{
		SetBit(LIGHT_FLOOR_IND1)
	} else{
		ClearBit(LIGHT_FLOOR_IND1)
	}
	if floor == 0x01{
		SetBit(LIGHT_FLOOR_IND2)
	} else{
		ClearBit(LIGHT_FLOOR_IND2)
	}
}



