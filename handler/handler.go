package handler

import (
	"fmt"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/utils"
)

func HandleEmail() {

	err, tasks := client.GetTasks()

	if err != nil {
		panic(err)
	}

	todayTasks := utils.GetTodayTasks(tasks)

	message := utils.BuildMessage(todayTasks)

	fmt.Println(message)

	// err = email.SendEmail("Tarefas do dia", message)

	if err != nil {
		panic(err)
	}
}
