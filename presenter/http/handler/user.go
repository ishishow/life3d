package handler

import (
	"encoding/json"
	"lifegame/presenter/dcontext"
	"lifegame/presenter/http/response"
	"lifegame/usecase"
	"log"
	"net/http"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return UserHandler{
		UserUseCase: uu,
	}
}

// HandleCreate ユーザを作成するHandler
func (h UserHandler) HandleCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody userCreateRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// データベースにユーザデータを登録する
		user, err := h.UserUseCase.Create(requestBody.Name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 生成した認証トークンを返却
		response.Success(writer, &userCreateResponse{Token: user.AuthToken})
	}
}

type userCreateRequest struct {
	Name string
}

type userCreateResponse struct {
	Token string
}

// HandleGet ユーザー取得処理
func (h UserHandler) HandleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		user, err := h.UserUseCase.Get(userID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, &userGetResponse{
			ID:   user.ID,
			Name: user.Name,
		})
	}
}

type userGetResponse struct {
	ID   string
	Name string
}