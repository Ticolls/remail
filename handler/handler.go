package handler

import (
	"fmt"
	"time"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/models"
	"github.com/Ticolls/remail/utils"
)

func Init() {
	for {
		curentTime := time.Now()

		if curentTime.Hour() == 7 && curentTime.Minute() == 0 {
			fmt.Println("São 7 horas.")

			EmailTaskHandler()
		}
	}
}

func EmailTaskHandler() {

	dailyTicker := time.NewTicker(24 * time.Hour)
	hourTicker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})

	var todayTasks *[]models.Task

	err, tasks := client.GetTasks()

	if err != nil {
		err = email.SendEmail("Erro recuperando as tarefas", err.Error())
		fmt.Println(err.Error())

	} else {
		todayTasks = utils.GetTodayTasks(tasks)

		err = sendTodayTasks(todayTasks)
		err = sendRememberTask(todayTasks)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	for {
		select {
		case <-dailyTicker.C:
			err, tasks := client.GetTasks()

			if err != nil {
				err = email.SendEmail("Erro recuperando as tarefas", err.Error())
				fmt.Println(err.Error())

			} else {
				todayTasks = utils.GetTodayTasks(tasks)
				err := sendTodayTasks(todayTasks)
				if err != nil {
					fmt.Println(err)
				}
			}

		case <-hourTicker.C:
			err := sendRememberTask(todayTasks)
			if err != nil {
				fmt.Println(err)
			}

		case <-quit:
			dailyTicker.Stop()
			hourTicker.Stop()
			return
		}
	}

}
