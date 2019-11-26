package database

type UserDAO interface {
	CreateUser(CreateUserOption) error
	GetUser(GetUserOption) (User, error)
}
