// TCPServer project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Printf("Receive from client | %s | %s | %s \n", rAddr.Network(), rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte(fmt.Sprintf("I have received :%s \n", string(buf[0:n]))))
		if err2 != nil {
			return
		}
	}
}

func main() {
	service := "127.0.0.1:10000"
	listener, err := net.Listen("tcp", service)
	fmt.Printf("Listen to %s \n", service)
	CheckErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
