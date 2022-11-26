package entity

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"updated_at,omitempty"`
}

type Page struct {
	Size      int
	Offset    int
	PageCount int
}
