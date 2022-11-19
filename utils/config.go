package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

var filePath string = path.Join(".", "config", "config.json")

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

func Decode() Config {
	var config Config
	json.Unmarshal(ReadConfig(), &config)
	return config
}

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

func createConfigDir() {
	_, err := os.Stat("config")
	if err != nil {
		os.Mkdir("config", 0775)
	}
}

func createConfigFile() {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defaultConfig := Config{
		Host:     "90.125.35.111",
		Port:     "8080",
		Pseudo:   "",
		Password: "",
	}
	content, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	f.Write(content)
}
