package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/models"
	"github.com/Ticolls/remail/utils"
)

func sendTodayTasks(tasks *[]models.Task) error {

	var message string
	var err error

	if len(*tasks) == 0 {
		message = "Sem tarefas para hoje!"
	} else {
		var emailTasks []models.EmailTask

		for _, task := range *tasks {

			_, project := client.GetProject(task.ProjectId)

			_, section := client.GetSection(task.SectionId)

			emailTask := models.NewEmailTask(&task, project, section)
			emailTasks = append(emailTasks, *emailTask)

		}

		message = utils.BuildDailyMessage(emailTasks)

		fmt.Println(message)
	}

	err = email.SendEmail("Tarefas do dia", message)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return errors.New("Erro mandando email.")
	}

	fmt.Println("Email diário enviado com sucesso!")

	return nil
}

func sendRememberTask(tasks *[]models.Task) error {

	currentTime := time.Now()
	hourNow := currentTime.Hour()
	minuteNow := currentTime.Minute()

	for _, task := range *tasks {

		if task.Due.Datetime != "" {
			taskTime := strings.Split(strings.Split(task.Due.Datetime, "T")[1], ":")

			hourString := taskTime[0]
			minuteString := taskTime[1]

			hour, err := strconv.Atoi(hourString)
			minute, err := strconv.Atoi(minuteString)

			if err != nil {
				return err
			}

			if (hour-hourNow) == 1 || (hour-hourNow) == 0 {
				var (
					message string
					subject string
				)

				_, project := client.GetProject(task.ProjectId)
				_, section := client.GetSection(task.SectionId)

				emailTask := models.NewEmailTask(&task, project, section)

				message = utils.BuildRememberMessage(*emailTask)

				if (hour-hourNow) == 1 && (minute-minuteNow) == 0 {
					subject = fmt.Sprintf("1 hora para a tarefa: %s", emailTask.Name)
				} else if (hour-hourNow) == 0 && minute > minuteNow {
					subject = fmt.Sprintf("%d minutos para a tarefa: %s", (minute - minuteNow), emailTask.Name)
				}

				err := email.SendEmail(subject, message)

				if err != nil {
					fmt.Printf("error: %s\n", err.Error())
					return errors.New("Erro mandando email.")
				}

				fmt.Println("remember email enviado com sucesso!")
			}
		}
	}

	return nil
}
