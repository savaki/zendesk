package main

import (
	"log"

	"github.com/savaki/zendesk"
)

func main() {
	client, err := zendesk.FromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	users, err := client.Users().List()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("found %d users\n", len(users))
}
