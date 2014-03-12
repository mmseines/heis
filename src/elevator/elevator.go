package elevator

import (
	"../order"
	"../driver"
	"fmt"
)

const (
	IDLE = 0
	OPEN = 1
	UP = 2
	DOWN = 3
	
	ON = 1
	OFF = 0
)

var (
	currentFloorTest, state, nextstate int
)


func ControlStateMachine() {
	switch state {
	
	case IDLE:
		if order.CheckCurrentFloor() {
			nextstate = IDLE
		} else if order.FindDirection() == 1 {
			nextstate = DOWN
		} else if order.FindDirection() == 0 {
			nextstate = UP
		}
		break
		
	case OPEN:
		// if true
		// timer ferdig
		nextstate = IDLE
		break
		
	case UP:
		if order.CheckCurrentFloor() {
			nextstate = OPEN
		} // else if "safety" {	
			//nextstate = IDLE
		//}
		break
		
	case DOWN:
		if order.CheckCurrentFloor() {
			nextstate = OPEN
		} // else if "safety" {	
			//nextstate = IDLE
		//}
		break
		
	default:
		break
	
	}
	
	if state != nextstate {
		fmt.Println(state, nextstate)
		order.PrintTable()
		fmt.Println(order.CheckCurrentFloor())
		switch nextstate {
	
		case IDLE:
			driver.ElevSetDoorOpenLamp(OFF)
			driver.ElevSetSpeed(0) // Maa haandtere braastopp-tingen
			break
			
		case OPEN:
			driver.ElevSetSpeed(0) // Maa haandtere braastopp-tingen
			driver.ElevSetDoorOpenLamp(ON)
			// start timer
			break
	
		case UP:
			driver.ElevSetDoorOpenLamp(OFF)
			driver.ElevSetSpeed(300) // verdi?
			break
	
		case DOWN:
			driver.ElevSetDoorOpenLamp(OFF)
			driver.ElevSetSpeed(-300) // verdi?
			break
	
		default:
			break
		}
	}

	state = nextstate

}

func Init(){
	state = IDLE
	nextstate = IDLE
}

