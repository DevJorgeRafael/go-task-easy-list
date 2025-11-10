package model

import "time"

type TaskPriority struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
}
