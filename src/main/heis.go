package main

import (
	"fmt"
	"../comm"
	//"../cost"
	"../elev"
	"net"
	//"time"
	//"os"
)



func main() {

	localaddr, err := net.ResolveUDPAddr("udp4", ":0")
	groupaddr, err := net.ResolveUDPAddr("udp4", "224.0.0.2:12000")
	//testaddr1, err := net.ResolveUDPAddr("udp4", "129.241.187.148:0")
	//testaddr2, err := net.ResolveUDPAddr("udp4", "129.241.187.147:0")
	comm.CheckError(err)

	localconn, err := net.ListenUDP("udp", localaddr)
	//groupconn, err := net.ListenUDP("udp", groupaddr)
	groupconn, err := net.ListenMulticastUDP("udp", nil, groupaddr)
	comm.CheckError(err)
	
	fmt.Println("Sockets created successfully")
	
	// Map functions-test
	//testmap := comm.NewPeermap()
	
	/*
	go comm.UpdatePeermap(testmap, groupconn)
	for {
		fmt.Println(comm.CheckPeerLife(*testmap, testaddr1))
		fmt.Println()
		fmt.Println(comm.CheckPeerLife(*testmap, testaddr1))
	}
	*/
	
	//Test of sending/recieving JSON
	testOrder := elev.NewOrder(2, 1)
	fmt.Println(testOrder)
	
	comm.CastData(testOrder, groupconn, localconn, groupaddr)
	

	localconn.Close()
	groupconn.Close()

	fmt.Println("End")
}
