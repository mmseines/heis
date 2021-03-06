package com

import (
	"../types"
	"net"
	"fmt"
	//"sync"
	"time"
	"encoding/gob"
	"os"
	//"../order"
)

var peerMap = NewPeerMap()
var OutputCh = make(chan types.Data)

// Create sockets and start go routines
func Run() {
	
	broadcastAddr := "129.241.187.255:12000" // For sanntidssalen
	//listenAddr := ":12000"

	//lAddr, err := net.ResolveUDPAddr("udp", listenAddr)
	bAddr, err := net.ResolveUDPAddr("udp", broadcastAddr)
	CheckError(err)

	//lConn, err := net.ListenUDP("udp", lAddr)
	bConn, err := net.DialUDP("udp", nil, bAddr)
	CheckError(err)
	
	fmt.Println("Sockets created successfully")

	// Global channels
	
	//OrderCh := make(chan []int)
	//TableCh := make(chan [][]int)
	//AuctionCh := make(chan int)
	OutputCh := make(chan types.Data)
	

	// Local channels
	peerCh := make(chan int)
	
	fmt.Println("Channels created succesfully")
	
	testData := types.Data{"cost", []int{1, 0, 1}, [][]int{}, 2, types.CART_ID, time.Now()}

	go CastData(bConn, OutputCh)
	go UpdatePeerMap(peerMap, peerCh)

	for {

		OutputCh <- testData
	}

	//go ChannelTester(peerch, orderch, tablech, aucch)
	//ReceiveData(lconn, peerch, orderch, tablech, aucch)
}


func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// Creates a map with peer IP as key and timer as element
func NewPeerMap() *types.PeerMap {
	return &types.PeerMap{M: make(map[int]time.Time)}
}

// Checks if peer address is in peer map and time difference is not > 1 sec
// New version using ID instead of peeraddr
func CheckPeerLife(p types.PeerMap, id int) bool {
	_, present := p.M[id]
	if present {
		tdiff := time.Since(p.M[id])
		return tdiff <= types.TIMEOUT
	}
	return false
}


// Updates peermap and sets discovery time from conn input
// New version using ID instead of IP
func UpdatePeerMap(p *types.PeerMap, peerCh chan int) {
	for {
		id := <- peerCh
		p.Mu.Lock()
		p.M[id] = time.Now()
		p.Mu.Unlock()
	}
}


// Listens and receives from connection in seperate go-routine
func ReceiveData(conn *net.UDPConn, peerCh chan int, orderch chan []int, tablech chan [][]int, aucch chan int) {
	var inc types.Data
	decoder := gob.NewDecoder(conn)
	for {
		err := decoder.Decode(&inc)
		
		fmt.Println(inc)
		CheckError(err)
		// update peermap
		peerCh <- inc.ID // c1
		
		if inc.Head == "order" {
			orderch <- inc.Order // c2
		} else if inc.Head == "table" {
			tablech <- inc.Table // c3
		} else if inc.Head == "cost" {
			aucch <- inc.Cost // c4
		}
		
		//fmt.Println(inc)
	}
}


func CastData(conn *net.UDPConn, OutputCh chan types.Data) {
	var d types.Data
	encoder := gob.NewEncoder(conn)
	for {
		d = <- OutputCh
		d.T = time.Now() // Sets timestamp on outgoing data
		err := encoder.Encode(d)
		CheckError(err)
		//fmt.Println(d)
	}
}

/*
func ReceiveTest(c *net.UDPConn, b []byte) {
	c.ReadFromUDP(b)
	fmt.Println(b)
}
*/

func ChannelTester(c1 chan int, c2 chan []int, c3 chan [][]int, c4 chan int) {
	for {
		select {
		case <-c1:
			fmt.Println("c1 read")
		case <-c2:
			fmt.Println("c2 read")
		case <-c3:
			fmt.Println("c3 read")
		case <-c4:
			fmt.Println("c4 read")
		}
	}
}


