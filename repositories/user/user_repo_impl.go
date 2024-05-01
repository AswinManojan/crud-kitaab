package repo

import (
	"context"
	"fmt"

	"github.com/sample-crud-app/repositories/user/models"
	"github.com/sample-crud-app/utils"
)

type RepoImpl struct {
}

// GetAllUsers implements repointer.RepoInter.
func (r *RepoImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := utils.DB.NewSelect().Model(&users).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}

// AddUser implements repointer.RepoInter.
func (r *RepoImpl) AddUser(user *models.User) (*models.User, error) {
	_, err := utils.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		fmt.Println(err, "- repo layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements repointer.RepoInter.
func (r *RepoImpl) GetUserByID(id int) (*models.User, error) {
	user := new(models.User)
	err := utils.DB.NewSelect().Model(user).Where("id=?", id).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

// GetUserByName implements repointer.RepoInter.
func (r *RepoImpl) GetUserByName(name string) (*models.User, error) {
	user := new(models.User)
	err := utils.DB.NewSelect().Model(user).Where("name=?", name).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func NewRepoImpl() *RepoImpl {
	return &RepoImpl{}
}
