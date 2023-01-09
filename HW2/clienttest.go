package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var nickName string
	fmt.Println("Введи свой никнейм: ")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())

	fmt.Scanf("%s\n", &nickName)
	conn.Write([]byte(nickName))

}
