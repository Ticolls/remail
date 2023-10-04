package main

import (
	"fmt"
	"log"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/handler"
)

func main() {

	err := email.Init()

	if err != nil {
		log.Fatalln(err)
		return
	}

	client.Init()

	fmt.Println("Everything ok with email and client services, waiting for ticks...")

	handler.Init()

}
