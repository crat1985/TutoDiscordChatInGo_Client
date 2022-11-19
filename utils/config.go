package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

func (c Config) PrintInfos() {
	fmt.Printf("Adresse : %s\nPort : %s\nPseudo : %s\nMot de passe : %s\n", c.Host, c.Port, c.Pseudo, c.Password)
}

func ReadConfig() []byte {
	createConfigDir()
	_, err := os.Stat(path.Join(".", "config", "config.json"))
	if err != nil {
		createConfigFile()
	}
	content, err := os.ReadFile(path.Join(".", "config", "config.json"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return content
}

func Decode() Config {
	var config Config
	json.Unmarshal(ReadConfig(), &config)
	return config
}

func createConfigDir() {
	_, err := os.Stat("config")
	if err != nil {
		os.Mkdir("config", 0775)
	}
}

func createConfigFile() {
	f, err := os.Create(path.Join(".", "config", "config.json"))
	if err != nil {
		panic(err)
	}
	defaultConfig := Config{
		Host:     "90.125.35.111",
		Port:     "8888",
		Pseudo:   "",
		Password: "",
	}
	content, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	f.Write(content)
}
