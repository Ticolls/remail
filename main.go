package main

import (
	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/handler"
)

func main() {

	email.Init()

	client.Init()

	handler.Init()

}
