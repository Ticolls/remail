package handler

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/models"
	"github.com/Ticolls/remail/utils"
)

func Init(tasks *[]models.Task) {
	for {
		curentTime := time.Now()

		if curentTime.Hour() == 18 && curentTime.Minute() == 39 {
			fmt.Println("São 7 horas.")
			EmailTaskHandler(tasks)
		}
	}
}

func EmailTaskHandler(tasks *[]models.Task) {

	dailyTicker := time.NewTicker(24 * time.Hour)
	hourTicker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})

	err := sendTodayTasks(tasks)
	// err = rememberTask(tasks)

	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		select {
		case <-dailyTicker.C:
			sendTodayTasks(tasks)

		case <-hourTicker.C:
			// rememberTask(tasks)
		case <-quit:
			dailyTicker.Stop()
			hourTicker.Stop()
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
		var emailTasks []models.EmailTask

		for _, task := range todayTasks {

			var (
				hour        string
				projectName string
				sectionName string
			)

			if task.Due.Datetime != "" {
				hour = strings.Split(task.Due.Datetime, "T")[1]
			}

			err, project := client.GetProject(task.ProjectId)
			if err == nil && project != nil {
				projectName = project.Name
			}

			err, section := client.GetSection(task.SectionId)
			if err == nil && section != nil {
				sectionName = section.Name
			}

			emailTasks = append(emailTasks, models.EmailTask{
				Name:        task.Content,
				Description: task.Description,
				Hour:        hour,
				ProjectName: projectName,
				SectionName: sectionName,
			})
		}

		message = utils.BuildMessage(emailTasks)
	}

	err = email.SendEmail("Tarefas do dia", message)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return errors.New("Erro mandando email.")
	}

	fmt.Println("Email enviado com sucesso!")

	return nil
}

// func rememberTask(tasks *[]models.Task) error {

// 	for _, task := range *tasks {

// 	}

// 	return nil
// }
