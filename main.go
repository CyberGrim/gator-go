package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/cybergrim/gator-go/internal/config"
	"github.com/cybergrim/gator-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	s := &state{db: database.New(db), Cfg: &cfg}

	cmds := commands{Cmd: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("Need to pass a command to run")
	}

	cmd := command{Name: arguments[1], Arguments: arguments[2:]}
	runError := cmds.run(s, cmd)
	if runError != nil {
		log.Fatal(runError)
	}
}
