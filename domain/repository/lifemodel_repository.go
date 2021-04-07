package repository

import "lifegame/domain/model"

type LifeModelRepository interface {
	Create() error
	Get(ID string) (*model.LifeModel, error)
}
