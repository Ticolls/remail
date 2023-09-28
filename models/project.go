package models

type Project struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	CommentCount   int64  `json:"comment_count"`
	Color          string `json:"color"`
	IsShared       bool   `json:"is_shared"`
	Order          int64  `json:"order"`
	IsFavorite     bool   `json:"is_favorite"`
	IsInboxProject bool   `json:"is_inbox_project"`
	IsTeamInbox    bool   `json:"is_team_inbox"`
	ViewStyle      string `json:"view_style"`
	Url            string `json:"url"`
	Parent_id      string `json:"parent_id"`
}
