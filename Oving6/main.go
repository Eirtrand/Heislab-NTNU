package main
	
import(

"fmt"
"time"
. "./network"
)

var b = 0
func main() {

var backup StatusMessage
backup.State = 1
for {
	elevMsg := ReadUDP("129.241.187.255:20012")
	if elevMsg.State == 1 {
		fmt.Println("I am backup")
		for(backup.State == 1){
			backup = ReadUDP("129.241.187.255:20012")
			if backup.State == 1 {b = backup.I}
		}
	} else{
		fmt.Println("I am primary")
		var elevMsg StatusMessage
		elevMsg.I = b
		fmt.Println("b er faktisk dene verdien naa:",b)
		
		for{
			elevMsg.I = elevMsg.I + 1
			SendUDP("129.241.187.255:20012", elevMsg)
			fmt.Println(elevMsg.I)
			time.Sleep(1*time.Millisecond)
		}
	}
}
}