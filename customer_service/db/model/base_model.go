package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type BaseModel struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}


func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UTC()
	m.ModifiedAt = time.Now().UTC()
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = time.Now().UTC()
	return
}