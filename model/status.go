package model

import "time"

type Status struct {
	ID       string    `json:"id" gorm:"index:id,unique;PRIMARY_KEY"`
	Online   bool      `json:"online"`
	LastSeen time.Time `json:"last_seen"`
}
