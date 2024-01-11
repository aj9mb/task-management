package model

type BoardUser struct {
	BoardId int64 `json:"boardId,omitempty"`
	UserId  int64 `json:"userId,omitempty"`
	AddedBy int64 `json:"addedBy,omitempty"`
}
