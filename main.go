package main


import (
	"log"
	"github.com/WoutHofstra/blogGator/internal/config"
	"os"
	_ "github.com/lib/pq"
	"github.com/WoutHofstra/blogGator/internal/database"
	"database/sql"
)

type state struct {
	db 	*database.Queries
	dbConn	*sql.DB
	cfg	*config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfgStruct := &state{}
	cfgStruct.cfg = &cfg

        dbUrl := cfg.DbURL
        db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

        cfgStruct.db = database.New(db)
	cfgStruct.dbConn = db

	cmdStruct := &commands{
		cmdNames: make(map[string]func(*state, command) error),
	}

	cmdStruct.register("login", handlerLogin)
	cmdStruct.register("register", handlerRegister)
	cmdStruct.register("reset", handlerReset)
	cmdStruct.register("users", handlerUsers)
	cmdStruct.register("agg", handlerAgg)
	cmdStruct.register("addfeed", middlewareLoggedIn(handlerFeed))
	cmdStruct.register("feeds", handlerGetFeeds)
	cmdStruct.register("follow", middlewareLoggedIn(handlerFollow))
	cmdStruct.register("following", middlewareLoggedIn(handlerFollowing))
	cmdStruct.register("unfollow", middlewareLoggedIn(handlerUnfollow))


	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments given")
	}
	cmdname := args[1]
	cmdArguments := args[2:]

	newCmdStruct := &command{
		name:		 cmdname,
		arguments:	 cmdArguments,
	}

	err = cmdStruct.run(cfgStruct, *newCmdStruct)
	if err != nil {
		log.Fatal(err)
	}
}
