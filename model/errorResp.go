package model

import "time"

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Url        string `json:"url"`
	Time       time.Time
}
