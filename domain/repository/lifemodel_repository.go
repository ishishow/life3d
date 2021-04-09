package repository

import "lifegame/domain/model"

type LifeModelRepository interface {
	Create(lifeModel *model.LifeModel) error
	Get(ID string) (*model.LifeModel, error)
	SetFavorite(ID string, userID string) error
	Ranking() error
}
