package model

type Board struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedBy int64  `json:"createdBy"`
}
