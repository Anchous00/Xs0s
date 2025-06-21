package server

import (
	"fmt"
	"net"
)

type Game struct {
	Current byte
	Field   [3][3]byte
}

var IsConnected bool
var IsServerRunning bool
var Conn net.Conn

func StartServer() {
	IsServerRunning = true

	go func() {
		listener, err := net.Listen("tcp", ":4545")

		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			for {
				if !IsServerRunning {
					listener.Close()
					return
				}
			}
		}()

		fmt.Println("Server is listening...")
		IsConnected = false
		Conn, err = listener.Accept()
		IsConnected = true
		if !IsServerRunning {
			IsConnected = false
		}
		if err != nil {
			fmt.Println(err)
			return
		}

	}()
}

func SendMove(code []byte) {
	Conn.Write(code)

}

func HandleCode(code []byte, Field [3][3]byte) [3][3]byte {
	char := code[0]
	fmt.Println(code)
	i := code[1]
	j := code[2]
	Field[i][j] = char
	return Field
}

func StartClient(ip string) bool {
	var err error
	Conn, err = net.Dial("tcp", ip+":4545")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func WaitMove(Field [3][3]byte) ([3][3]byte, byte) {
	code := make([]byte, 3)
	_, err := Conn.Read(code[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("recieved code", code)
	Field = HandleCode(code, Field)
	if code[0] == 'X' {
		code[0] = '0'
	} else {
		code[0] = 'X'
	}

	return Field, code[0]
}
