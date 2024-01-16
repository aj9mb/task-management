package model

import "time"

type TaskComment struct {
	Id       int64      `json:"id,omitempty"`
	UserId   int64      `json:"userId,omitempty"`
	UserName *string    `json:"name,omitempty"`
	Comment  *string    `json:"comment,omitempty"`
	TaskId   int64      `json:"taskId,omitempty"`
	CreateDt *time.Time `json:"createdDate,omitempty"`
	LastUpt  *time.Time `json:"lastUpdate,omitempty"`
}
