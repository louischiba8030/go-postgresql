package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
//	"fmt"
)

type UUIDBaseModel struct {
	ID uuid.UUID `gorm:"primaryKey;unique;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `"gorm:index"`
}

// This functions are called before creating Post

type Post struct {
	Name string
	Age uint
	Bloodtype string
	Origin string
}

func (p *Post) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}
