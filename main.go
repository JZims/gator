package main

import (
	"database/sql"
	config "gator/internal"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	configFile, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", configFile.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &configFile,
	}

	cmds := &commands{
		options: make(map[string]func(*state, command) error),
	}

	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
