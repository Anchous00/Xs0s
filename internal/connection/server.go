package server

import (
	"fmt"
	"net"
)

type Game struct {
	Current byte
	Field   [3][3]byte
}

var ip net.Addr
var Conn net.Conn

func StartServer() {

	go func() {
		listener, err := net.Listen("tcp", ":4545")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Server is listening...")

		Conn, err = listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("New connection from ", Conn.RemoteAddr())
		ip = Conn.RemoteAddr()

	}()
}

func SendMove(code []byte) {
	fmt.Println("sent code", code)
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

func StartClient(ip string) {
	var err error
	fmt.Println("voshlo")
	Conn, err = net.Dial("tcp", ip+":4545")
	fmt.Println(Conn.RemoteAddr())
	if err != nil {
		fmt.Println(err)
		return
	}

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
