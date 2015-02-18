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

var buttonMap = map[int]int{
		FLOOR_COMMAND1: 0,
		FLOOR_COMMAND2: 1,
		FLOOR_COMMAND3: 2,
		FLOOR_COMMAND4: 3,
		FLOOR_UP1:      4,
		FLOOR_UP2:      5,
		FLOOR_UP3:      6,
		FLOOR_DOWN2:    7,
		FLOOR_DOWN3:    8,
		FLOOR_DOWN4:    9,
	}
	func clearButtonLight(int buttonKey){
		for i
	}
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
	
	if (Read_bit(SENSOR_FLOOR1)) {
		return 0;
	}
	else if (Read_bit(SENSOR_FLOOR2)) {
		return 1;
	}
	else if (Read_bit(SENSOR_FLOOR3)) {
		return 2;
	}
	else if (Read_bit(SENSOR_FLOOR4)) {
		return 3;
	}
	else {
		return -1;
	}
}

func elev_set_floor_indicator(int floor) {
    // Binary encoding. One light must always be on.
    if (floor >= 0 && floor < N_FLOORS){
	    if (floor & 0x02){
	        Set_bit(LIGHT_FLOOR_IND1);
		}
	    else{
	        Clear_bit(LIGHT_FLOOR_IND1);
		}
	    if (floor & 0x01){
	        Set_bit(LIGHT_FLOOR_IND2);
		}
	    else{
	        Clear_bit(LIGHT_FLOOR_IND2);
		}
    }
	else {
		fmt.Printf("elev_set_floor_indicator: Elevator out of range.")
	}
}












