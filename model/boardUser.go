package model

type BoardUser struct {
	BoardId int64 `json:"boardId"`
	UserId  int64 `json:"userId"`
	AddedBy int64 `json:"addedBy"`
}
