package driver

import (
	"fmt"
	"elevproj/io.go"


)
type lampChannelMatrix= [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

	

type Order struct {
	Floor int
	Direction OrderDirection
}
	
type OrderDirection int
const (
	ORDER_UP OrderDirection = iota
	ORDER_DOWN
	ORDER_INTERNAL
)
const 
const N_BUTTONS 3
const N_FLOORS 4

var buttonMap = map[int]Order{
	FLOOR_COMMAND1: {1, ORDER_INTERNAL},
	FLOOR_COMMAND2: {2, ORDER_INTERNAL},
	FLOOR_COMMAND3: {3, ORDER_INTERNAL},
	FLOOR_COMMAND4: {4, ORDER_INTERNAL},
	FLOOR_UP1:      {1, ORDER_UP},
	FLOOR_UP2:      {2, ORDER_UP},
	FLOOR_UP3:      {3, ORDER_UP},
	FLOOR_DOWN2:    {2, ORDER_DOWN},
	FLOOR_DOWN3:    {3, ORDER_DOWN},
	FLOOR_DOWN4:    {4, ORDER_DOWN},
}

var lightMap = map[int]int{
	SENSOR1: 1,
	SENSOR2: 2,
	SENSOR3: 3,
	SENSOR4: 4,
}

func clearButtonLight(int buttonKey){
	for i
}

func Init() () {
	if !IoInit() {
		fmt.Printf("IO not initiated\n")
		return 0	
	}
	else {
		fmt.Printf("IO initiated\n")
		
	}
	for i := 0; i < N_FLOORS;i++ {
		if i != 0 {
			SetButtonLamp(ORDER_DOWN, i, 0)
		}
		if I != N_FLOORS-1 {
			SetButtonLamp(ORDER_UP, i, 0)
		}
		SetButtonLamp(ORDER_INTERNAL, i, 0)
		
	}
	SetStopLamp(0)
	SetDoorOpenLamp(0)
	SetFloorIndicator(0)
}

func SetStopLamp(i int) {
	if i == 1 {
		IoSetBit(LIGHT_STOP)
	}
	else {
		IoClearBit(LIGHT_STOP)
	}
}

func SetDoorOpenLamp(i int) {
	if i == 1 {
		IoSetBit(DOOR_OPEN)
	}
	else{
		IoClearBit(DOOR_OPEN)
	}
}

func SetFloorIndicator(i int) {

}

func SetButtonLamp(button OrderDirection, floor int, val int) {

	if floor >= 0 && floor < N_FLOORS {
		if dir == "up" && floor == 3 {
			fmt.Printf("The current direction and floor does not exist (up and 4)")
			return
		}
		if dir == "down" && floor == 0 {
			fmt.Printf("The current direction and floor does not exist (down and 0)")
			return
		}
		if val == 1 {
			IoSetBit(lampChannelArray[floor][int(button)])		
		}	
		else {
			IoClearBit(lampChannelArray[floor][int(button)]
		}	
	else {
		fmt.Printf("Floor and direction is out of bounds")
		return				


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


func clearLights() {

	ClearLight(buttonMap[1])	
	ClearLight(buttonMap[2])	
	ClearLight(buttonMap[3])
	ClearLight(buttonMap[4])
	ClearLight(buttonMap[5])
	ClearLight(buttonMap[6])
	ClearLight(buttonMap[7])
	ClearLight(buttonMap[8])
	ClearLight(buttonMap[9])

