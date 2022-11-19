package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/RIC217/TutoDiscordChatInGo_Client/utils"
)

var conn net.Conn
var err error
var address = "90.125.35.111"
var port = "8080"
var pseudo = ""
var password = ""

func askInfos() {
	fmt.Print("Adresse du serveur : ")
	fmt.Scanln(&address)
	fmt.Print("Port : ")
	fmt.Scanln(&port)
	fmt.Print("Pseudo : ")
	fmt.Scanln(&pseudo)
	fmt.Print("Mot de passe : ")
	fmt.Scanln(&password)
}

func sendPseudo() {
	for {
		config := utils.Decode()
		address = "90.125.35.111"
		port = "8080"
		pseudo = ""
		password = ""
		var ask = true
		if config.Pseudo != "" {
			for {
				fmt.Print("Se connecter avec les infos enregistrées (o pour oui, n pour non et ? pour plus d'infos) ? ")
				var response string
				fmt.Scanln(&response)
				response = strings.ToLower(response)
				if response == "o" {
					address = config.Host
					port = config.Port
					pseudo = config.Pseudo
					password = config.Password
					ask = false
					break
				}
				if response == "?" {
					config.PrintInfos()
					continue
				}
				if response == "n" {
					break
				}
			}
		}
		if ask {
			askInfos()
		}
		response := make([]byte, 1024)
		log.Printf("Connecting to %s:%s...\n", address, port)
		conn, err = net.Dial("tcp", address+":"+port)
		if err != nil {
			log.Print(err)
			continue
		}
		log.Println("Connected !")
		conn.Write([]byte(string(pseudo) + "\n" + string(password)))
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
	}
}

func main() {
	sendPseudo()
	go sendMessage(conn)
	listenForMessages(conn)
}

func sendMessage(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var slice []byte = make([]byte, 1024)
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
