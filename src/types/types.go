package types

import (
    "time"
    "sync"
)


const (
	CART_ID = 0
	N_FLOORS = 4
	N_BUTTONS = 4
	
	// Global timeout const
	TIMEOUT = 1 * time.Second
)



type (
    ElevButtonTypeT int
    
    LocalTable [][]int
    GlobalTable [][]int

    
    // Map for storing addresses of peers in group
    PeerMap struct {
	    Mu sync.Mutex
	    M map[int]time.Time
    }

    // Struct sent over network
    Data struct {
	    Head string
	    Order []int
	    Table [][]int
	    Cost int
	    ID int
	    T time.Time
    }

)

func NewLocalTable() LocalTable {
	t := make([][]int, N_FLOORS)
	for i := range t {
		t[i] = make([]int, 3)
		for j := range t[i] {
			t[i][j] = 0
		}
	}
	return t
}

func NewGlobalTable() GlobalTable {
	t := make([][]int, N_FLOORS)
	for i := range t {
		t[i] = make([]int, 4)
		for j := range t[i] {
			t[i][j] = 0
		}
	}
	return t
}



