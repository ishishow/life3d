package handler

import (
	"lifegame/presenter/dcontext"
	"lifegame/presenter/http/response"
	"lifegame/usecase"
	"log"
	"net/http"
)

type LifeModelHandler struct {
	LifeModelUseCase usecase.LifeModelUseCase
}

func NewLifeModelHandler(lu usecase.LifeModelUseCase) LifeModelHandler {
	return LifeModelHandler{
		LifeModelUseCase: lu,
	}
}

func (lh LifeModelHandler) HandleCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		if err := lh.LifeModelUseCase.Create(userID, make([]int, 5, 5)); err != nil {
			log.Printf("%v", err)
			response.InternalServerError(writer, "error")
		}
		response.Success(writer, &createLifeModelResponse{})
	}
}

func (lh LifeModelHandler) HandleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		lifemodel, err = lh.LifeModelUseCase.Get(userID)
		if err != nil {
			log.Printf("%v", err)
			response.InternalServerError(writer, "error")
		}

		response.Success(writer, &getLifeModelResponse{
			LifeModel: lifemodel,
		})
	}
}

func (lh LifeModelHandler) HandleRanking() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//response.Success(writer, &rankingLifeModelResponse{
		//	LifeModelList:
		//})
	}
}

type createLifeModelRequest struct {
	UserID string
	Map    []int
}

type createLifeModelResponse struct{}

type getLifeModelRequest struct {
	ID string
}

type getLifeModelResponse struct {
	LifeModel lifeModel
}

type rankingLifeModelRequest struct{}

type rankingLifeModelResponse struct {
	LifeModelList []lifeModel
}

type lifeModel struct {
	ID       string
	Name     string
	Map      []int
	UserName string
	Favorite int
}
