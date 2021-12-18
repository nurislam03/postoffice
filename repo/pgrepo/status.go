package pgrepo

import (
	"github.com/nurislam03/postoffice/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
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

func (b *Status) Upsert(sts *model.Status) error {
	return b.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"id", "online", "last_seen"}),
	}).Create(sts).Error
}

func (b *Status) Expire(sts *model.Status) error {
	Now := time.Now()
	lastTime := Now.Add(-30*time.Second)
	return b.DB.Where("last_seen < ?", lastTime).Delete(sts).Error
}