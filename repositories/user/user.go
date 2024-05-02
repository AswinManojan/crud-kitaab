package user

import (
	"context"
	"fmt"

	"github.com/sample-crud-app/repositories/user/models"
	"github.com/sample-crud-app/utils"
)

type UserRepository struct{}

// GetAllUsers implements repointer.RepoInter.
func (r *UserRepository) QueryAll() ([]models.User, error) {
	var users []models.User
	err := utils.DB.NewRaw("select * from users").Scan(context.Background(), &users)
	return users, err
}

// AddUser implements repointer.RepoInter.
func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	_, err := utils.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		fmt.Println(err, "- repo layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements repointer.RepoInter.
func (r *UserRepository) QueryByID(id int) (*models.User, error) {
	user := new(models.User)
	err := utils.DB.NewRaw("select * from users where user_id=?", id).Scan(context.Background(), user)
	return user, err
}

func (r *UserRepository) Delete(id int) (bool, error) {
	res, err := utils.DB.NewDelete().Model(&models.User{}).Where("user_id=?", id).Exec(context.Background())
	fmt.Println(res)
	return true, err
}

func (r *UserRepository) Update(id int, user *models.User) (*models.User, error) {
	_, err := utils.DB.NewUpdate().Model(&models.User{}).Set("name=?", user.Name).Set("height=?", user.Height).Set("birth_date=?", user.BirthDate).Where("user_id=?", id).Exec(context.Background())
	return user, err
}
