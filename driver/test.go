
package main
import (
"elev"
"fmt"
"time"
)

func main(){
	fmt.Printf("Current floor: ", elev.GetFloorSensorSignal())
	elev.SetLightFloorIndicator(elev.GetFloorSensorSignal)
	fmt.Printf("Drive!")	
	elev.SetSpeed(20)
	time.Sleep(2*time.Second)
	fmt.Printf("Stop!")
	elev.SetSpeed(0)

	fmt.Printf("Button lamps on")
	elev.SetButtonLamp(ORDER_UP, 0, 1) 
	elev.SetButtonLamp(ORDER_UP, 1, 1) 
	elev.SetButtonLamp(ORDER_UP, 2, 1) 

	elev.SetButtonLamp(ORDER_DOWN, 1, 1) 
	elev.SetButtonLamp(ORDER_DOWN, 2, 1) 
	elev.SetButtonLamp(ORDER_DOWN, 3, 1) 

	fmt.Printf("Button lamps off")
	elev.SetButtonLamp(ORDER_UP, 0, 0) 
	elev.SetButtonLamp(ORDER_UP, 1, 0) 
	elev.SetButtonLamp(ORDER_UP, 2, 0) 

	elev.SetButtonLamp(ORDER_DOWN, 1, 0) 
	elev.SetButtonLamp(ORDER_DOWN, 2, 0) 
	elev.SetButtonLamp(ORDER_DOWN, 3, 0) 


	fmt.Printf("Open door")
	elev.SetOpenDoor(1)
	fmt.Printf("Stop lamp")
	elev.SetStopLamp(1)
	fmt.Printf("Door open")
	elev.SetDoorOpenLamp(1)
	fmt.Printf("Obstruction signal")
	elev.GetObstructionSignal()
	fmt.Printf("Stop signal")
	elev.GetStopSignal()
	fmt.Printf("Floor sensor signal")
	elev.GetFloorSensorSignal()
	fmt.Printf("Button signal")
	elev.GetButtonSignal()
	 


}
