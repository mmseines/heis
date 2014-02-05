package comm

import (
	"net"
	"fmt"
	"sync"
	"time"
	//"encoding/json"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}


// leser alle incoming packets med ReadFromUDP og sorterer dem etter innholder/lengde.
// bestemme om en selv er doed ved aa sjekke timeout paa ReadFromMulticast

// Map for storing addresses of peers in group
type peermap struct {
	mu sync.Mutex
	m map[string]*time.Timer
}

// Creates a map with peer IP as key and timer as element
func NewPeermap() *peermap {
	return &peermap{m: make(map[string]*time.Timer)}
}

// Checks if peer address is in peer map and has timer value >0
// unfinished
func CheckPeerLife(p peermap, addr net.Addr) bool {
	fmt.Println(addr)
	peeraddr, _, err := net.SplitHostPort(addr.String())
	CheckError(err)
	_, present := p.m[peeraddr]

	if present {
		tValue := <- p.m[peeraddr].C
		return !tValue.IsZero()
	}
	return false
}

// Updates peermap and sets timer from conn input
func UpdatePeermap(p *peermap, conn *net.UDPConn) {
	var buf [1024]byte
	_, addr, err := conn.ReadFrom(buf[:])
	CheckError(err)
	peeraddr, _, err := net.SplitHostPort(addr.String())
	CheckError(err)
	p.mu.Lock()
	_, present := p.m[peeraddr]
	if present {
		p.m[peeraddr].Reset(5 * time.Second)
	} else {
		p.m[peeraddr] = time.NewTimer(5 * time.Second)	
	}
	p.mu.Unlock()
}

// Adds an address (with port number) to a peerlist
/*
func AddToPeermap(p *peermap, a net.Addr) {
	p.mu.Lock()
	p.m[a.String()] = true
	p.mu.Unlock()
}
*/


// Creates multicast-listener for UDP packages in a multicast group. Error check for network operations -- rename to JoinGroup()?

// This is in main() for now
func CreateSocket() {
}

// Receive data from multicast socket. Returns number of bytes read and the return address of the packet. Can be made to timeout and return an error after a fixed time limit; see SetDeadline and SetReadDeadline.
func ReceiveData(conn *net.UDPConn) ([]byte) {
	fmt.Println("ReceiveData begin")
	data := make([]byte, 256)
	_, _, err := conn.ReadFromUDP(data)
	CheckError(err)
	fmt.Println("read", string(data))
	fmt.Println("ReceiveData end")
	return data
}

func CastData(data []byte, conn, lconn *net.UDPConn, gaddr *net.UDPAddr) {
	fmt.Println("CastData begin")
	_, _ = lconn.WriteToUDP(data, gaddr)
	fmt.Println("CastData end")
}


// Updates an array/slice with the address of the clients sending out multicast packet and whether or not it's alive. If clientlist does not exist -> create. If client already in list -> do nothing. If client reappears after timeout -> do not make a new entry to the list. Instead set it's status as alive.
//func UpdateClientList()

//func SortMulticastPacket(



