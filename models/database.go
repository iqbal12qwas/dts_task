package models

import (
    "time"
)

type Task struct {
    ID         uint      `json:"id" gorm:"primary_key"`
	IdProfile int    `json:"id_profile" validate:"required"`
    AssignedTo string    `json:"assigned_to" validate:"required"`
    Task       string    `json:"task" validate:"required"` 
    Deadline   time.Time `json:"deadline" validate:"required"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Profile struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Name string `json:"name" validate:"required"`
    Birth   time.Time `json:"birth" validate:"required"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName ...
func (Task) TableName() string {
	return "tasks"
}

func (Profile) TableName() string {
	return "profiles"
}
