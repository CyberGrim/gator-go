package main

import (
	"fmt"
	"log"

	"github.com/cybergrim/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("Grim")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
