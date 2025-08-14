package main


import (
	"fmt"
	"log"
	"github.com/WoutHofstra/blogGator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("Wout")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
