package user

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(userRequest Register) (User, error)
	Update(ID int, userRequest Register) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(userRequest Register) (User, error) {

	user := User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Update(ID int, userRequest Register) (User, error) {
	user, err := s.repository.FindByID(ID)

	user.Username = userRequest.Username
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	newUser, err := s.repository.Update(user)
	return newUser, err
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	delUser, err := s.repository.Delete(user)
	return delUser, err
}
