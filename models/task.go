package models

import "strings"

type due struct {
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime"`
	String      string `json:"string"`
	Timezone    string `json:"timezone"`
}

type duration struct {
	Amount int64  `json:"amount"`
	Unit   string `json:"unit"`
}

type Task struct {
	CreatorId    string   `json:"creator_id"`
	CreatedAt    string   `json:"created_at"`
	AssigneeId   string   `json:"assignee_id"`
	AssignerId   string   `json:"assigner_id"`
	CommentCount int64    `json:"comment_count"`
	IsCompleted  bool     `json:"is_completed"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	Due          due      `json:"due"`
	Duration     duration `json:"duration"`
	Id           string   `json:"id"`
	Labels       []string `json:"labels"`
	Order        int64    `json:"order"`
	Priority     int64    `json:"priority"`
	ProjectId    string   `json:"project_id"`
	SectionId    string   `json:"section_id"`
	ParentId     string   `json:"parent_id"`
	Url          string   `json:"url"`
}

type EmailTask struct {
	Name        string
	Description string
	Hour        string
	ProjectName string
	SectionName string
}

func NewEmailTask(task *Task, project *Project, section *Section) *EmailTask {
	var (
		hour        string
		projectName string
		sectionName string
	)

	if task.Due.Datetime != "" {
		hour = strings.Split(task.Due.Datetime, "T")[1]
	}

	if project != nil {
		projectName = project.Name
	}

	if section != nil {
		sectionName = section.Name
	}

	emailTask := EmailTask{
		Name:        task.Content,
		Description: task.Description,
		Hour:        hour,
		ProjectName: projectName,
		SectionName: sectionName,
	}

	return &emailTask
}
