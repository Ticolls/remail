package utils

import (
	"strings"
	"time"

	"github.com/Ticolls/remail/models"
)

func GetTodayTasks(allTasks []models.Task) []models.Task {

	var todayTasks []models.Task

	currentTime := time.Now().String()
	date := strings.Split(currentTime, " ")[0]

	for _, task := range allTasks {

		taskDate := task.Due.Date
		if taskDate == date {
			todayTasks = append(todayTasks, task)
		}
	}

	return todayTasks
}

func getHour(datetime string) string {
	return strings.Split(datetime, "T")[1]
}
