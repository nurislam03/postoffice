package repo

import "github.com/nurislam03/postoffice/model"

type StatusRepo interface {
	Create(status *model.Status) error
	//Update(status *model.Status) error
	//Delete(status *model.Status) error
}
