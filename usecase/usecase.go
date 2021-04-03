package usecase

type UserUseCase interface {
	Create(name string) (authToken string, err error)
	Get(userID string) (user *domain.User, err error)
}

type userUseCase struct {
	repository repository.Userrepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: userRepo,
	}
}
func (uu userUseCase) Create(name string) (string, error) {
	return "authToken", nil
}

func (uu userUseCase) Get(userID string) (*domain.User, error) {
	return nil, nil
}
