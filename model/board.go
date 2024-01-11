package model

type Board struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedBy int64  `json:"createdBy,omitempty"`
}
