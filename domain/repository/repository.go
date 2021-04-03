package repository

type UserRepository interface {
	Create(name string) (authToken string, err error)
	Get(userID string) (user *domain.User, err error)
}
