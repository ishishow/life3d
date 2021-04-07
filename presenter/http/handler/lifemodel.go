package handler

import "lifegame/usecase"

type LifeModelHandler struct {
	LifeModelUseCase usecase.LifeModelUseCase
}

func NewLifeModelHandler(lu usecase.LifeModelUseCase) LifeModelHandler {
	return LifeModelHandler{
		LifeModelUseCase: lu,
	}
}

func (lh lifeModelHandler) {

}
