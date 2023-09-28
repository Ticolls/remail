package models

type Section struct {
	Id        string `json:"id"`
	ProjectId string `json:"project_id"`
	Order     int64  `json:"order"`
	Name      string `json:"name"`
}
