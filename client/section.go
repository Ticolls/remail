package client

import (
	"fmt"

	"github.com/Ticolls/remail/models"
)

func GetSection(id string) (error, *models.Section) {

	url := fmt.Sprintf("https://api.todoist.com/rest/v2/sections/%s", id)

	var section models.Section

	err := getJson(url, &section)

	if err != nil {
		fmt.Printf("Error getting section: %s\n", err.Error())
		return err, nil
	}

	return nil, &section
}
