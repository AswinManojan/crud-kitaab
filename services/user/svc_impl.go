package svc

import (
	"fmt"

	repo "github.com/sample-crud-app/repositories/user"
	"github.com/sample-crud-app/repositories/user/models"
)

type SVCImpl struct {
	Repo *repo.RepoImpl
}

// GetAllUsers implements svcinter.SVCInter.
func (s *SVCImpl) GetAllUsers() ([]*models.User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return users, nil
}

// AddUser implements svcinter.SVCInter.
func (s *SVCImpl) CreateUser(user *models.User) (*models.User, error) {
	user, err := s.Repo.AddUser(user)
	if err != nil {
		fmt.Println(err, "- in svc layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements svcinter.SVCInter.
func (s *SVCImpl) GetUserByID(id int) (*models.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}

// GetUserByName implements svcinter.SVCInter.
func (s *SVCImpl) GetUserByName(name string) (*models.User, error) {
	user, err := s.Repo.GetUserByName(name)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}

func NewSVCImpl(repo *repo.RepoImpl) *SVCImpl {
	return &SVCImpl{Repo: repo}
}
