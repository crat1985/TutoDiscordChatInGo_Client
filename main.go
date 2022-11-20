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

// Connexion
var conn net.Conn
var err error

// Adresse du serveur
var address = "90.125.35.111"

// Port du serveur
var port = "8888"

// Pseudo
var pseudo = ""

// Mot de passe
var password = ""
var ask = true

// Structure de configuration
var config utils.Config

// Demande à l'utilisateur des informations par le biais de la fonction Scanln du package fmt
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

// Exécutée si la configuration est valide
func validConfig() {
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

// Envoie le pseudo renseigné au serveur et vérifie que la réponse du serveur est correcte
func sendPseudo() {
	for {
		config = utils.Decode()
		address = "90.125.35.111"
		port = "8888"
		pseudo = "admin"
		password = "password"
		if config.IsValid() {
			validConfig()
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
		for {
			fmt.Print("Enregistrer cette connexion comme la connexion par défaut (o/n) ? ")
			var response string
			fmt.Scanln(&response)
			response = strings.ToLower(response)
			if response == "o" {
				log.Println("Enregistrement en cours...")
				err := utils.Encode(utils.Config{Host: address, Port: port, Pseudo: pseudo, Password: password})
				if err != nil {
					log.Print(err)
					break
				}
				log.Println("Enregistrement effectué avec succès !")
				break
			}
			if response == "n" {
				break
			}
		}
		log.Println("Bienvenue sur le chat !")
		break
	}
}

// Fonction principale
func main() {
	sendPseudo()
	go sendMessage(conn)
	listenForMessages(conn)
}

// Attend que l'utilisateur écrive un message dans la console puis l'envoie au serveur
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

// Ecoute les message envoyés par le serveur puis les affiche dans la console
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
