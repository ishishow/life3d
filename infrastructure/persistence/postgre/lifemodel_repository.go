package postgre

import (
	"lifegame/domain/model"
	"lifegame/domain/repository"
)

type lifeModelRepositoryImpl struct {
	SQLHandler
}

func NewLifeModelRepositoryImpl(sqlHandler SQLHandler) repository.LifeModelRepository {
	return &lifeModelRepositoryImpl{
		SQLHandler: sqlHandler,
	}
}

func (lri lifeModelRepositoryImpl) Create() error {
	return nil
}

func (lri lifeModelRepositoryImpl) Get(ID string) (*model.LifeModel, error) {
	return nil, nil
}
