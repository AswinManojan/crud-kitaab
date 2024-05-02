package user

import (
	"errors"
	"fmt"

	repo "github.com/sample-crud-app/repositories/user"
	"github.com/sample-crud-app/repositories/user/models"
)

type SVCImpl struct {
	Repo *repo.RepoImpl
}

// GetAllUsers implements svcinter.SVCInter.
func (s *SVCImpl) QueryAll() ([]models.User, error) {
	users, err := s.Repo.QueryAll()
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return users, nil
}

// AddUser implements svcinter.SVCInter.
func (s *SVCImpl) Create(user *models.User) (*models.User, error) {
	user, err := s.Repo.Create(user)
	if err != nil {
		fmt.Println(err, "- in svc layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements svcinter.SVCInter.
func (s *SVCImpl) QueryByID(id int) (*models.User, error) {
	user, err := s.Repo.QueryByID(id)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}
func (s *SVCImpl) Delete(id int) (bool, error) {
	if _, err := s.Repo.QueryByID(id); err != nil {
		return false, errors.New("error fetching the user")
	}
	res, err := s.Repo.Delete(id)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return res, err
	}
	return res, nil
}

// GetUserByName implements svcinter.SVCInter.
func (s *SVCImpl) QueryByName(name string) (*models.User, error) {
	user, err := s.Repo.QueryByName(name)
	if err != nil {
		fmt.Println(err, "- svc layer")
		return nil, err
	}
	return user, nil
}
func (s *SVCImpl) Update(id int, user *models.User) (*models.User, error) {
	user.UserID = uint(id)

	usr, err := s.Repo.Update(id, user)
	if err != nil {
		fmt.Println("Error updating user in svc layer")
		return nil, err
	}
	return usr, nil
}

func NewSVCImpl(repo *repo.RepoImpl) *SVCImpl {
	return &SVCImpl{Repo: repo}
}
