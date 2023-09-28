package utils

import (
	"fmt"
	"strings"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/models"
)

func BuildMessage(tasks []models.Task) (error, string) {
	var message string

	for i, task := range tasks {
		message = message + fmt.Sprintf("tarefa %d: %s\n", i+1, task.Content)
		if task.Description != "" {
			message = message + fmt.Sprintf("Descrição: %s\n", task.Description)
		}
		if task.Due.Datetime != "" {
			hour := strings.Split(task.Due.Datetime, "T")[1]
			message = message + fmt.Sprintf("Horário: %s\n", hour)
		}
		err, project := client.GetProject(task.ProjectId)

		if err != nil {
			return err, ""
		}

		message = message + fmt.Sprintf("Project: %s\n", project.Name)
		message = message + "\n"
	}

	return nil, message
}
