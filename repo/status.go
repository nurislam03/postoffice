package repo

import "github.com/nurislam03/postoffice/model"

type StatusRepo interface {
	Create(status *model.Status) error
	Upsert(status *model.Status) error
	Expire(status *model.Status) error
}
