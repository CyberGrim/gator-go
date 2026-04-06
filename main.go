package main

import (
	"log"
	"os"

	"github.com/cybergrim/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{Cfg: &cfg}

	cmds := commands{Cmd: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("Need to pass a command to run")
	}

	cmd := command{Name: arguments[1], Arguments: arguments[2:]}
	err2 := cmds.run(s, cmd)
	if err2 != nil {
		log.Fatal(err2)
	}
}
