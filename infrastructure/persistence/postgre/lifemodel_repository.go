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

func (lri *lifeModelRepositoryImpl) Create(userID string, lifeMap []int) error {

	return nil
}

func (lri *lifeModelRepositoryImpl) Get(ID string) (*model.LifeModel, error) {
	//row := lri.SQLHandler.Conn.QueryRow("SELECT * FROM users WHERE id = $1", ID)
	return nil, nil
}
