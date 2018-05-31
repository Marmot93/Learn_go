// UDPServer project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, rAddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Printf("Receive from client | %s | %s | %s \n", rAddr.Network(), rAddr.String(), string(buf[0:n]))
		_, err2 := conn.WriteToUDP([]byte("Welcome udp!"), rAddr)
		if err2 != nil {
			return
		}
	}
}
func main() {
	service := "127.0.0.1:10000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	fmt.Printf("Listen to %s \n", service)
	checkErr(err)
	for {
		handleClient(conn)
	}
}
