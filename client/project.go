package client

import (
	"fmt"

	"github.com/Ticolls/remail/models"
)

func GetProject(id string) (error, *models.Project) {

	url := fmt.Sprintf("https://api.todoist.com/rest/v2/projects/%s", id)

	var project models.Project

	err := getJson(url, &project)

	if err != nil {
		fmt.Printf("Error getting project: %s\n", err.Error())
		return err, nil
	}

	return nil, &project
}
