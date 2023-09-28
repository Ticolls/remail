package handler

import (
	"fmt"
	"time"

	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/models"
	"github.com/Ticolls/remail/utils"
)

func Init(tasks *[]models.Task) {
	for {
		curentTime := time.Now()

		if curentTime.Hour() == 13 && curentTime.Minute() == 45 {
			fmt.Println("São 13 horas.")
			EmailTaskHandler(tasks)
		}
	}
}

func EmailTaskHandler(tasks *[]models.Task) {

	ticker := time.NewTicker(24 * time.Hour)
	quit := make(chan struct{})

	sendTodayTasks(tasks)

	for {
		select {
		case <-ticker.C:
			sendTodayTasks(tasks)
		case <-quit:
			ticker.Stop()
			return
		}
	}

}

func sendTodayTasks(tasks *[]models.Task) {
	todayTasks := utils.GetTodayTasks(tasks)

	var message string

	if len(todayTasks) == 0 {
		message = "Sem tarefas para hoje!"
	} else {
		message = utils.BuildMessage(todayTasks)
	}

	err := email.SendEmail("Tarefas do dia", message)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	fmt.Println("Email enviado com sucesso!")
}
