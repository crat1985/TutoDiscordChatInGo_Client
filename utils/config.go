package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

// Chemin vers le fichier de configuration
var filePath string = path.Join(".", "config", "config.json")

// Structure de configuration
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// Afficher les informations de configuration
func (c Config) PrintInfos() {
	fmt.Printf("Adresse : %s\nPort : %s\nPseudo : %s\nMot de passe : %s\n", c.Host, c.Port, c.Pseudo, c.Password)
}

// Renvoie vrai si la configuration est valide, faux sinon
func (c Config) IsValid() bool {
	if c.Host != "" && c.Port != "" && c.Pseudo != "" && c.Password != "" {
		return true
	}
	return false
}

// Lit le fichier de configuration et renvoie son contenu
func ReadConfig() []byte {
	createConfigDir()
	_, err := os.Stat(filePath)
	if err != nil {
		createConfigFile()
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return content
}

// Décode du texte sous forme JSON vers une structure de type Config
func Decode() Config {
	var config Config
	json.Unmarshal(ReadConfig(), &config)
	return config
}

// Encode une structure de type Config en JSON et écrit le résultat dans le fichier de configuration
func Encode(c Config) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	return err
}

// Crée le dossier contenant le fichier de configuration
func createConfigDir() {
	_, err := os.Stat("config")
	if err != nil {
		os.Mkdir("config", 0775)
	}
}

// Crée le fichier de configuration avec les informations par défaut
func createConfigFile() {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defaultConfig := Config{
		Host:     "90.125.35.111",
		Port:     "8888",
		Pseudo:   "example",
		Password: "example",
	}
	content, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	f.Write(content)
}
