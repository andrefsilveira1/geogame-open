package models

import "time"

type Player struct {
	Id        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Age       uint      `json:"age,omitempty"`
	Score     int       `json:"score,omitempty"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}
