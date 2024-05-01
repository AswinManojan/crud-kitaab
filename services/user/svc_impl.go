package user

import (
	"fmt"

	repo "github.com/sample-crud-app/repositories/user"
	"github.com/sample-crud-app/repositories/user/models"
)

type SVCImpl struct {
	Repo *repo.RepoImpl
}

// GetAllUsers implements svcinter.SVCInter.
func (s *SVCImpl) GetAll() ([]*models.User, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return users, nil
}

// AddUser implements svcinter.SVCInter.
func (s *SVCImpl) Create(user *models.User) (*models.User, error) {
	user, err := s.Repo.Add(user)
	if err != nil {
		fmt.Println(err, "- in svc layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements svcinter.SVCInter.
func (s *SVCImpl) GetByID(id int) (*models.User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}

// GetUserByName implements svcinter.SVCInter.
func (s *SVCImpl) GetByName(name string) (*models.User, error) {
	user, err := s.Repo.GetByName(name)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}

func NewSVCImpl(repo *repo.RepoImpl) *SVCImpl {
	return &SVCImpl{Repo: repo}
}
