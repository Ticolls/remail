package utils

import (
	"strings"
	"time"

	"github.com/Ticolls/remail/models"
)

func GetTodayTasks(tasks *[]models.Task) []models.Task {

	var todayTasks []models.Task

	currentTime := time.Now().String()
	date := strings.Split(currentTime, " ")[0]

	for _, task := range *tasks {

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

func getTasksWithHour(tasks []models.Task) []models.Task {
	var tasksWithHour []models.Task

	for _, task := range tasks {
		if getHour(task.Due.Datetime) != "" {
			tasksWithHour = append(tasksWithHour, task)
		}
	}

	return tasksWithHour
}
