package repository

import "lifegame/domain/model"

type LifeModelRepository interface {
	Create(userID string, lifeMap []int) error
	Get(ID string) (*model.LifeModel, error)
}
