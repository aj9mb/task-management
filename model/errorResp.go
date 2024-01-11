package model

import "time"

type ErrorResponse struct {
	StatusCode int       `json:"status,omitempty"`
	Message    string    `json:"message,omitempty"`
	Url        string    `json:"url,omitempty"`
	Time       time.Time `json:"time,omitempty"`
}
