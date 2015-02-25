package elevator

import("fmt"
	"time"
	. "../driver"
) 

const n_BUTTONS int = 3
const n_FLOORS int = 4



var elev elevator


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




func InitializeElevator() int {

	if !Init() {
		fmt.Println("init failed")
		return 0
	}
	for GetFloorSensorSignal() == -1{
		SetMotorDirection(DOWN)
	}

	//externalButtonChannel := make(chan int)
	//internalButtonChannel := make(chan int)

	elev.externalButtons = make([]int, NUMBER_OF_EXT_BUTTONS)
	elev.internalButtons = make([]int, NUMBER_OF_INT_BUTTONS)

	go PollAllInputs()
	//go InputHandler(internalButtonChannel, externalButtonChannel)
	
	

	SetMotorDirection(NONE)
	return 1
}


func PollAllInputs() {

	for{
		for i := 0; i<NUMBER_OF_INT_BUTTONS; i++{
				if GetButtonSignal(internal_button_array[i]) {
					elev.internalButtons[i] = 1
				}
			}

		for i := 0; i<NUMBER_OF_EXT_BUTTONS; i++{
			if GetButtonSignal(external_button_array[i]) {
				elev.externalButtons[i] = 1
			}	
			
		}	
	time.Sleep(time.Millisecond * 18)
	}
}


func PrintElev() {
	fmt.Println("STATE: ", elev.state)



	fmt.Println("EXTERNAL BUTTONS")
	if(elev.externalButtons[5] == 1){
		fmt.Println("	[O]")
	}else {
		fmt.Println("	[ ]")
    }

	if(elev.externalButtons[4] == 1){
		fmt.Print("	[O]")
	}else {
		fmt.Print("	[ ]")
    }

	if(elev.externalButtons[2] == 1){
		fmt.Println("[O]")
	}else {
		fmt.Println("[ ]")
    }

	if(elev.externalButtons[3] == 1){
		fmt.Print("	[O]")
	}else {
		fmt.Print("	[ ]")
    }

	if(elev.externalButtons[1] == 1){
		fmt.Println("[O]")
	}else {
		fmt.Println("[ ]")
    }

	if(elev.externalButtons[0] == 1){
		fmt.Println("	[O]")
	}else {
		fmt.Println("	[ ]")
    }


    fmt.Println("INTERNAL BUTTONS")
    if(elev.internalButtons[0] == 1){
    		fmt.Print("   	[O]")
	}else {
		fmt.Print("   	[ ]")
    }

    if(elev.internalButtons[1] == 1){
    		fmt.Print("[O]")
	}else {
		fmt.Print("[ ]")
    }
    if(elev.internalButtons[2] == 1){
    		fmt.Print("[O]")
	}else {
		fmt.Print("[ ]")
    }

    if(elev.internalButtons[3] == 1){
    		fmt.Print("[O]")
	}else {
		fmt.Print("[ ]")
    }

	
}




