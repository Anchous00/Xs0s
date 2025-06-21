package server

import (
	"fmt"
	"net"
)

type Game struct {
	Current byte
	Field   [3][3]byte
}

var Conn net.Conn

func StartServer() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Ошибка при получении интерфейсов:", err)
		return
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("Ошибка при получении адресов для %s: %s\n", i.Name, err)
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Printf("Интерфейс: %s, IP: %s\n", i.Name, v.IP.String())
			case *net.IPAddr:
				fmt.Printf("Интерфейс: %s, IP: %s\n", i.Name, v.IP.String())
			}
		}
	}

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
