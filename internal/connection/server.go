package server

import (
	"fmt"
	"log"
	"net"
)

var ip string

var Gamefield [3][3]string

func StartServer() {
	go func() {
		listener, err := net.Listen("tcp", ":4545")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer listener.Close()
		fmt.Println("Server is listening...")
		conn, err := listener.Accept()
		defer conn.Close()
		for {

			if err != nil {
				fmt.Println(err)
				return
			}
			for {
				var buffer [3]byte // 1KB buffer to read data
				_, err := conn.Read(buffer[:])
				if err != nil {
					log.Printf("err while reading from conn: %v, exiting ...", err)
					return
				}
				i := buffer[1]
				j := buffer[2]
				player := buffer[0]
				Gamefield[i][j] = string(player)
			}
		}
	}()
}

func ConnectToServer() {
	ip = "127.0.0.1"
}

func SendMove(code []byte) {
	conn, err := net.Dial("tcp", ip+":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write(code)
}
