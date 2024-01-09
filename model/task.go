package model

type Task struct {
	Id       int64  `json:"id"`
	BoardId  int64  `json:"boardId"`
	AddedBy  int64  `json:"addedBy"`
	Assignee int64  `json:"assigned_to"`
	TaskDesc string `json:"taskDesc"`
}
