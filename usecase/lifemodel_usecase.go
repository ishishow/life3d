package usecase

import (
	"lifegame/domain/model"
	"lifegame/domain/repository"
	"log"
)

type LifeModelUseCase interface {
	Create(userID string, lifeMap []int) error
	Get(ID string) (*model.LifeModel, error)
	Ranking() ([]*model.LifeModel, error)
}

type lifeModelUseCase struct {
	lifeModelRepository repository.LifeModelRepository
	userRepository      repository.UserRepository
}

func NewLifeModelUseCase(lifeModelRepo repository.LifeModelRepository,
	userRepo repository.UserRepository) LifeModelUseCase {
	return &lifeModelUseCase{
		lifeModelRepository: lifeModelRepo, userRepository: userRepo,
	}
}

func (lu *lifeModelUseCase) Create(userID string, lifeMap []int) error {
	log.Println(userID, lifeMap)
	if err := lu.lifeModelRepository.Create(userID, lifeMap); err != nil {
		return err
	}
	return nil
}

func (lu *lifeModelUseCase) Get(ID string) (*model.LifeModel, error) {
	lifeModel, err := lu.lifeModelRepository.Get(ID)
	if err != nil {
		return lifeModel, err
	}

	return lu.fillUserDetails(lifeModel)
}

func (lu *lifeModelUseCase) Ranking() ([]*model.LifeModel, error) {
	return nil, nil
}

func (lu *lifeModelUseCase) fillUserDetails(lifeModel *model.LifeModel) (*model.LifeModel, error) {
	return lifeModel, nil
}
