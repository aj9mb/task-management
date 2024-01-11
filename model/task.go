package model

import "time"

type Task struct {
	Id           int64      `json:"id,omitempty"`
	BoardId      int64      `json:"boardId,omitempty"`
	AddedBy      int64      `json:"addedBy,omitempty"`
	Assignee     int64      `json:"assignedTo,omitempty"`
	TaskDesc     string     `json:"taskDesc,omitempty"`
	Status       bool       `json:"status,omitempty"`
	AddedByName  *string    `json:"addedByName,omitempty"`
	AssigneeName *string    `json:"assigneeName,omitempty"`
	CreateDt     *time.Time `json:"createdDate,omitempty"`
	LastUpt      *time.Time `json:"lastUpdate,omitempty"`
}
