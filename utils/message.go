package utils

import (
	"fmt"

	"github.com/Ticolls/remail/models"
)

func BuildDailyMessage(tasks []models.EmailTask) string {

	var message string

	for i, task := range tasks {

		message = message + fmt.Sprintf("tarefa %d: %s\n", i+1, task.Name)

		if task.Description != "" {
			message = message + fmt.Sprintf("Descrição: %s\n", task.Description)
		}
		if task.Hour != "" {
			message = message + fmt.Sprintf("Horário: %s\n", task.Hour)
		}
		if task.ProjectName != "" {
			message = message + fmt.Sprintf("Project: %s\n", task.ProjectName)
		}
		if task.SectionName != "" {
			message = message + fmt.Sprintf("Section: %s\n", task.SectionName)
		}

		message = message + "\n"
	}

	return message
}

func BuildRememberMessage(task models.EmailTask) string {
	var message string

	message = message + fmt.Sprintf("tarefa: %s\n", task.Name)

	if task.Description != "" {
		message = message + fmt.Sprintf("Descrição: %s\n", task.Description)
	}
	if task.Hour != "" {
		message = message + fmt.Sprintf("Horário: %s\n", task.Hour)
	}
	if task.ProjectName != "" {
		message = message + fmt.Sprintf("Project: %s\n", task.ProjectName)
	}
	if task.SectionName != "" {
		message = message + fmt.Sprintf("Section: %s\n", task.SectionName)
	}

	return message
}
