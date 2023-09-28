package handler

import (
	"fmt"
	"time"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/utils"
)

func HandleDailyEmail() {

	ticker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			curentTime := time.Now()

			if curentTime.Hour() == 8 {
				err, tasks := client.GetTasks()

				if err != nil {
					panic(err)
				}

				todayTasks := utils.GetTodayTasks(tasks)

				var message string

				if len(todayTasks) == 0 {
					message = "Sem tarefas para hoje!"
				} else {
					message = utils.BuildMessage(todayTasks)
				}

				fmt.Println(message)

				err = email.SendEmail("Tarefas do dia", message)

				if err != nil {
					panic(err)
				}
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}

}
