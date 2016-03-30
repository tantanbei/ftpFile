package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	b := make([]byte, 15000000)
	n, err := conn.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("read tcp bytes", n)
	b = b[:n]

	f, err := os.OpenFile("chexiang.sql", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}

	f.Write(b)
}
