package user

type UserRepository interface {
	
}

func NewUserRepository() UserRepository {
	return &Repository{}
}