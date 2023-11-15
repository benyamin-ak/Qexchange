package authentication

import "time"

type UserManagementService interface {
	Register(RegisterUserRequest) RegisterUserResponse
	Login(LoginUserRequest) LoginUserResponse
}

type UserRepository interface {
	Create(UserCreationRepository) error

	//returned string is password field value
	FetchUserByUsername(username string) (UserRepositoryResponseData, string, error)
}

type UserSessionStore interface {
	//returned string is user session id
	//time.Time use to set exipration time for user session, if you want a session with out expiration time pass 0
	Store(UserRepositoryResponseData, time.Time) (string, error)

	//string input is session id
	Expire(string) error
}
