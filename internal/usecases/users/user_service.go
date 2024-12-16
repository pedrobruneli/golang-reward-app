package users

type UserService struct {
	UserRepository *UserRepositoryImpl
}

func NewUserService(ur *UserRepositoryImpl) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) GetUsers() ([]User, error) {
	users, err := us.UserRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) CreateUser(user *User) error {
	err := us.UserRepository.CreateUser(user)
	return err
}

func (us *UserService) GetUserByUsername(username string) (*User, error) {
	user, err := us.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) DeleteUser(username string) error {
	err := us.UserRepository.DeleteUser(username)
	return err
}
