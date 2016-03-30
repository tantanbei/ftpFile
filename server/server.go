package main

import (
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("chexiang.sql", os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 15000000)

	n, err := f.Read(b)
	if err != nil || n > 15000000 {
		panic(err)
	}

	conn.Write(b[:n])
}
