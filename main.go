package main

import (
	"GoCart/config"
	"GoCart/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
