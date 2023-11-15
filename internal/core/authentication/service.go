package authentication

import (
	"errors"
	"time"
)

var passwordMismatched = errors.New("Wrong Password")

func NewUserService(userRepo UserRepository, sessionStore UserSessionStore) *userService {
	return &userService{
		UserRepo:     userRepo,
		SessionStore: sessionStore,
	}
}

func (srv *userService) Register(userData RegisterUserRequest) RegisterUserResponse {
	userRepoData := UserCreationRepository{
		RegisterUserRequest: userData,
		AuthenticationLevel: "basic",
		AccountStatus:       "active",
		accountType:         "simple",
	}

	err := srv.UserRepo.Create(userRepoData)
	if err != nil {
		return RegisterUserResponse{
			Status: false,
			Error:  err,
		}
	}
	return RegisterUserResponse{
		Status: true,
		Error:  nil,
	}
}

func (srv *userService) Login(loginData LoginUserRequest) LoginUserResponse {
	userData, password, err := srv.UserRepo.FetchUserByUsername(loginData.Username)
	if err != nil {
		return LoginUserResponse{
			Status:    false,
			SessionId: "",
			Error:     err,
		}
	}

	if !srv.checkPassword(loginData.Password, password) {
		return LoginUserResponse{
			Status:    false,
			SessionId: "",
			Error:     passwordMismatched,
		}
	}

	sessionId, err := srv.SessionStore.Store(userData, time.Now().Add(time.Minute*time.Duration(30)))
	if err != nil {
		return LoginUserResponse{
			Status:    false,
			SessionId: "",
			Error:     passwordMismatched,
		}
	}

	return LoginUserResponse{
		Status:    true,
		SessionId: sessionId,
		Error:     nil,
	}
}

func (srv *userService) checkPassword(enteredPassword string, realPassword string) bool {

	return true
}
