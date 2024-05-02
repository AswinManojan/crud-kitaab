package user

import (
	"github.com/sample-crud-app/repositories/user"
	"github.com/sample-crud-app/repositories/user/models"
)

type UserService struct{}

var userRepository *user.UserRepository

// GetAllUsers implements svcinter.SVCInter.
func (s *UserService) QueryAll() ([]models.User, error) {
	return userRepository.QueryAll()
}

// AddUser implements svcinter.SVCInter.
func (s *UserService) Create(user *models.User) (*models.User, error) {
	return userRepository.Create(user)
}

// GetUserByID implements svcinter.SVCInter.
func (s *UserService) QueryByID(id int) (*models.User, error) {
	return userRepository.QueryByID(id)
}

func (s *UserService) Delete(id int) (bool, error) {
	return userRepository.Delete(id)
}

func (s *UserService) Update(id int, user *models.User) (*models.User, error) {
	user.UserID = uint(id)
	return userRepository.Update(id, user)
}
