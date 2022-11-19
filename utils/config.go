package utils

import (
	"encoding/json"
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

func ReadConfig() []byte {
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
