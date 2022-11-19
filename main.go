package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var conn net.Conn
var err error

func sendPseudo() {
	response := make([]byte, 1024)
	log.Println("Connecting to 90.125.35.111:8080...")
	conn, err = net.Dial("tcp", "90.125.35.111:8080")
	if err != nil {
		panic(err)
	}
	log.Println("Connected !")

	for {
		fmt.Print("Pseudo : ")
		var pseudo []byte
		fmt.Scanln(&pseudo)
		conn.Write(pseudo)
		n, err := conn.Read(response)
		if err != nil {
			log.Println(err)
			continue
		}
		if string(response[:n]) != "pseudook" {
			log.Println(string(response[:n]))
			continue
		}
		log.Println("Bienvenue sur le chat !")
		break
	}

}

func main() {
	sendPseudo()
	go sendMessage(conn)
	listenForMessages(conn)
}

func sendMessage(conn net.Conn) {
	// for {
	// 	fmt.Println("Entre ton message :")
	// 	_, err := fmt.Scanln(&slice)
	// 	if err != nil {
	// 		// log.Print(err)
	// 		// continue
	// 		panic(err)
	// 	}
	// 	conn.Write(slice)
	// }
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var slice []byte = make([]byte, 1024)
		fmt.Println("Entre ton message :")
		for scanner.Scan() {
			bytes := scanner.Bytes()
			slice = bytes
			break
		}
		conn.Write(slice)
	}
}

func listenForMessages(conn net.Conn) {
	sliceMessage := make([]byte, 1024)
	var stringMessage string
	for {
		n, err := conn.Read(sliceMessage)
		if err != nil {
			log.Println("Lost connection to server !")
			return
		}
		stringMessage = string(sliceMessage[:n])
		splitedMessage := strings.Split(stringMessage, "\n")
		log.Println(splitedMessage[0] + ": " + splitedMessage[1])
	}
}
