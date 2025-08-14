package main


import (
	"log"
	"github.com/WoutHofstra/blogGator/internal/config"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfgStruct := &state{}

	cfgStruct.config = &cfg

	cmdStruct := &commands{
		cmdNames: make(map[string]func(*state, command) error),
	}

	cmdStruct.register("login", handlerLogin)

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
