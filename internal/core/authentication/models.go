package authentication

type userService struct {
	UserRepo     UserRepository
	SessionStore UserSessionStore
}

type RegisterUserRequest struct {
	FirstName    string
	LastName     string
	Age          int
	NationalCode string
	PhoneNumber  string
	Email        string
	UserName     string
	password     string
}

type RegisterUserResponse struct {
	Status bool
	Error  error
}

type LoginUserRequest struct {
	Username string
	Password string
}

type LoginUserResponse struct {
	Status    bool
	SessionId string
	Error     error
}

type UserCreationRepository struct {
	RegisterUserRequest
	AuthenticationLevel string
	AccountStatus       string
	accountType         string
}

type UserRepositoryResponseData struct {
	id                  uint
	AuthenticationLevel string
	AccountStatus       string
	AccountType         string
}
