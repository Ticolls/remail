package client

import (
	"fmt"

	"github.com/Ticolls/remail/models"
)

func GetTasks() {
	url := "https://api.todoist.com/rest/v2/tasks"

	var tasks []models.Task

	err := getJson(url, &tasks)

	if err != nil {
		fmt.Printf("Error getting tasks: %s\n", err.Error())
		return
	}

	fmt.Print(tasks)

}
