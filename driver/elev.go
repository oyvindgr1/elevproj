package driver

import (
	"fmt"
	"elevproj/io.go"


)

type Order struct {
	Floor int
	Direction 
}
	

const N_BUTTONS 3


func Init() () {
	if !IoInit() {
		fmt.Printf("IO initiated\n")
	}
	else {
		fmt.Printf("IO not initiated\n")
	}
	//Clear the lights!
	

}

func setDoorOpenLamp(int i) {
	if  i == 1 {
		Set_bit(LIGHT_DOOR_OPEN)
	}
	else {
		Clear_bit(LIGHT_DOOR_OPEN)
	}
}

//kommentar

func getObstructionSignal() bool {
	return Read_bit(OBSTRUCTION)
}

func getStopSignal() bool {
	return Read_bit(STOP)
}

func setStopLamp(int i) {
	Set_bit(LIGHT_STOP)
}

func getFloorSensor() int {
	
