package usecase

import (
	"lifegame/domain/model"
	"lifegame/domain/repository"
)

type LifeModelUseCase interface {
	Create() error
	Get(ID string) (*model.LifeModel, error)
}

type lifeModelUseCase struct {
	repository repository.LifeModelRepository
}

func NewLifeModelUseCase(lifeModelRepo repository.LifeModelRepository) LifeModelUseCase {
	return &lifeModelUseCase{
		repository: lifeModelRepo,
	}
}
func (lu *lifeModelUseCase) Create() error {
	return nil
}

func (lu *lifeModelUseCase) Get(ID string) (*model.LifeModel, error) {
	return nil, nil
}
