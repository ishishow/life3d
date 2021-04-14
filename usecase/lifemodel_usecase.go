package usecase

import (
	"lifegame/domain/model"
	"lifegame/domain/repository"

	"github.com/google/uuid"
)

type LifeModelUseCase interface {
	Create(lifeModel *model.LifeModel) error
	Get(ID string) (*model.LifeModel, error)
	Ranking() ([]*model.LifeModel, error)
	SetFavorite(ID string, userID string) error
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

func (lu *lifeModelUseCase) Create(lifeModel *model.LifeModel) error {
	ID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	lifeModel.ID = ID.String()
	return lu.lifeModelRepository.Create(lifeModel)
}

func (lu *lifeModelUseCase) Get(ID string) (*model.LifeModel, error) {
	lifeModel, err := lu.lifeModelRepository.Get(ID)
	if err != nil {
		return lifeModel, err
	}
	lifeModel.Favorite, err = lu.lifeModelRepository.GetFavoriteCount(ID)
	if err != nil {
		return lifeModel, err
	}
	return lu.fillUserDetails(lifeModel)
}

func (lu *lifeModelUseCase) SetFavorite(ID string, userID string) error {
	return lu.lifeModelRepository.SetFavorite(ID, userID)
}

func (lu *lifeModelUseCase) Ranking() ([]*model.LifeModel, error) {
	//lu.lifeModelRepository.
	return nil, nil
}

func (lu *lifeModelUseCase) fillUserDetails(lifeModel *model.LifeModel) (*model.LifeModel, error) {
	user, err := lu.userRepository.Get(lifeModel.User.ID)
	if err != nil {
		return nil, err
	}
	lifeModel.User = user
	return lifeModel, err
}
