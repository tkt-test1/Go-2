package models

import "gorm.io/gorm"

type TaskHistory struct {
	gorm.Model
	TaskID    uint   `json:"task_id"`
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}
