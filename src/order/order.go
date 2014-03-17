package order

import (
    "../types"
	"../driver"
	"fmt"
	"os"
	//"time"
)

const (
    UP = 0
	DOWN = 1
	INTERNAL = 2
)

//channel som henger sammen med lyssetting
var C1 = make(chan bool)

var (
	LocalOrders types.LocalTable
	GlobalOrders types.GlobalTable
)
    

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// INTERNAL maa erstattes, vurder assert
// Vurder navn paa denne, denne i en go routine?
func UpdateLocalTable(lt types.LocalTable, C1 chan bool) {
	for i := range lt {
		if lt[i][INTERNAL] != 1 {
			lt[i][INTERNAL] = driver.ElevGetButtonSignal(INTERNAL, i)
		}
	}
	C1 <- true
}

// INTERNAL maa erstattes, vurder assert
// Vurder navn paa denne
func RemoveOrder(lt types.LocalTable) {
	floor := driver.ElevGetFloorSensorSignal() 
	dir := driver.ElevGetDirection()
	if floor != -1 && floor < types.N_FLOORS {
		lt[floor][INTERNAL] = 0
		lt[floor][dir] = 0
	}
}

// Vurder assert, tar ikke hensyn til retning
func CheckCurrentFloor(lt types.LocalTable) bool {
	currentFloor := driver.ElevGetFloorSensorSignal()
	if currentFloor != -1 && currentFloor < types.N_FLOORS {
		return (lt[currentFloor][0] == 1 || lt[currentFloor][1] == 1 || lt[currentFloor][2] == 1)
	}
	return false
}

// For enkel, returnerer bare den foerste ordren den finner. Kan gjoeres om til aa returnere flere verdier
func CheckAllFloors(lt types.LocalTable) int {
	currentFloor := driver.ElevGetFloorSensorSignal()
	for floor := range lt {
		if floor != currentFloor {
			for i := 0; i < len(lt [floor]); i++ {
				if lt[floor][i] == 1 {
					return floor
				}
			}
		}
	}
	return -1
}


func FindDirection() int {
	var diff int
	currentFloor := driver.ElevGetFloorSensorSignal()
	if currentFloor != -1 && currentFloor < types.N_FLOORS && CheckAllFloors(LocalOrders) != -1 {
		diff = currentFloor - CheckAllFloors(LocalOrders)
	}
	if diff > 0 {
		return 1
	} else if diff < 0 {
		return 0
	} else {
		return -1
	}
}

func Init() {
	LocalOrders = types.NewLocalTable()
	GlobalOrders = types.NewGlobalTable()
}

func PrintTable(){
	fmt.Println(LocalOrders)
}


func SetLights(lt types.LocalTable, C1 chan bool){
    for{
        <- C1
        //time.Sleep(10 * time.Millisecond)
        for floor := range lt {
        	for i := 0; i < len(lt [floor]); i++ {
        	    driver.ElevSetLights(floor, i, lt[floor][i])
        	}
        }
    }
}
















