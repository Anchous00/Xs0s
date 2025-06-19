package server

import (
	"fmt"
	"net"
)

var ip = "192.168.1.164"

var Gamefield [3][3]string
var current string

func Init() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			Gamefield[i][j] = ""
		}
	}
	current = "X"
}

func StartServer() {
	Init()
	message := make([]byte, 3)
	message[0] = '0'
	message[1] = '1'
	message[2] = '2'
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Read(message[:])
		go HandleSignal(message)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				conn.Write([]byte(Gamefield[i][j]))
			}
		}

		conn.Close()
	}
}

func HandleSignal(message []byte) {
	player := message[0]
	i := message[1]
	j := message[2]
	if current == "X" {
		current = "0"
	}
	if current == "0" {
		current = "X"
	}
	Gamefield[i][j] = string(player)
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
