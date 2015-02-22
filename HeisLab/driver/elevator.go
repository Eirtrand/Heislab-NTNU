package driver

import("fmt"
	"time"
) 

const n_BUTTONS int = 3
const n_FLOORS int = 4

type Direction int
type State int
var elev elevator


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
}


func InitializeElevator() int {

	if !Init() {
		fmt.Println("init failed")
		return 0
	}
	externalButtonChannel := make(chan int)
	internalButtonChannel := make(chan int)
	floorSensorChannel := make(chan int)

	elev.externalButtons = make([]int, NUMBER_OF_EXT_BUTTONS)
	elev.internalButtons = make([]int, NUMBER_OF_INT_BUTTONS)

	go PollAllInputs(floorSensorChannel, externalButtonChannel, internalButtonChannel)
	go InputHandler(floorSensorChannel, externalButtonChannel, internalButtonChannel)
	
	for GetFloorSensorSignal() == 0{
		SetMotorDirection(DOWN)
	}

	SetMotorDirection(NONE)
	return 1
}


func PollAllInputs(floorSensorChannel, internalButtonChannel, externalButtonChannel chan int) {
	for{
		for i := 0; i<NUMBER_OF_EXT_BUTTONS; i++{
			if GetButtonSignal(external_button_array[i]) {
				externalButtonChannel <- external_button_array[i]
				fmt.Println(external_button_array[i])
			}
		}
		for i := 0; i<NUMBER_OF_INT_BUTTONS; i++{
			if GetButtonSignal(internal_button_array[i]) {
				internalButtonChannel <- internal_button_array[i]
				fmt.Println(internal_button_array[i])
			}
		}
		if GetFloorSensorSignal() != 0{
			floorSensorChannel <- GetFloorSensorSignal()
		}
		time.Sleep(time.Millisecond * 18)
	}	
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
		return 0
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

func InputHandler(floorSensorChannel, externalButtonChannel, internalButtonChannel chan int){
	for{
		select{
			case <-internalButtonChannel:
				elev.internalButtons[1] = <- internalButtonChannel
				fmt.Println("MOTTATT ORDRE:", elev.internalButtons[1])
	
			case <-externalButtonChannel:
				elev.externalButtons[1] = <- externalButtonChannel
				fmt.Println("MOTTATT ORDRE:", elev.externalButtons[1])
			case elev.currentFloor <- floorSensorChannel:
				if elev.currentFloor != elev.prevFloor && elev.currentFloor != 0{
					
				}
				elev.currentFloor = <- floorSensorChannel
				fmt.Println("currentFloor:", elev.currentFloor)
		}
	}
}

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