package client

import (
	"fmt"

	"github.com/Ticolls/remail/models"
)

func GetTasks() (error, *[]models.Task) {
	url := "https://api.todoist.com/rest/v2/tasks"

	var tasks []models.Task

	err := getJson(url, &tasks)

	if err != nil {
		fmt.Printf("Error getting tasks: %s\n", err.Error())
		return err, nil
	}

	return nil, &tasks
}
