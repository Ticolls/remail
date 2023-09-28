package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/models"
	"github.com/Ticolls/remail/utils"
)

func Init(tasks *[]models.Task) {
	for {
		curentTime := time.Now()

		if curentTime.Hour() == 14 && curentTime.Minute() == 18 {
			fmt.Println("São 14 horas.")
			EmailTaskHandler(tasks)
		}
	}
}

func EmailTaskHandler(tasks *[]models.Task) {

	ticker := time.NewTicker(24 * time.Hour)
	quit := make(chan struct{})

	err := sendTodayTasks(tasks)

	if err != nil {
		fmt.Println(err.Error())
	}

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

func sendTodayTasks(tasks *[]models.Task) error {
	todayTasks := utils.GetTodayTasks(tasks)

	var message string
	var err error

	if len(todayTasks) == 0 {
		message = "Sem tarefas para hoje!"
	} else {
		err, message = utils.BuildMessage(todayTasks)

		if err != nil {
			return errors.New("Erro construindo mensagem.")
		}
	}

	err = email.SendEmail("Tarefas do dia", message)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return errors.New("Erro mandando email.")
	}

	fmt.Println("Email enviado com sucesso!")

	return nil
}
