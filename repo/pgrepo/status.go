package pgrepo

import (
	"github.com/nurislam03/postoffice/model"
	"gorm.io/gorm"
)

type Status struct {
	*gorm.DB
}

func NewStatus(d *gorm.DB) *Status {
	return &Status{
		DB: d,
	}
}

func (b *Status) Create(sts *model.Status) error {
	return b.DB.Create(sts).Error
}
