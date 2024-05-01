package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sample-crud-app/repositories/user/models"
	"github.com/sample-crud-app/utils"
)

type RepoImpl struct {
}

// GetAllUsers implements repointer.RepoInter.
func (r *RepoImpl) GetAll() ([]models.User, error) {
	var users []models.User
	// err := utils.DB.NewSelect().Model(&users).Scan(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	rows, _ := utils.DB.Query("select * from users")
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Height, &user.BirthDate); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no users to fetch")
			}
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println(rows)
	return users, nil
}

// AddUser implements repointer.RepoInter.
func (r *RepoImpl) Add(user *models.User) (*models.User, error) {
	_, err := utils.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		fmt.Println(err, "- repo layer")
		return nil, err
	}
	return user, nil
}

// GetUserByID implements repointer.RepoInter.
func (r *RepoImpl) GetByID(id int) (*models.User, error) {
	user := new(models.User)
	// err := utils.DB.NewSelect().Model(user).Where("id=?", id).Scan(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// return user, nil
	row := utils.DB.QueryRow("select * from users where id=?", id)
	if err := row.Scan(&user.UserID, &user.Name, &user.Height, &user.BirthDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no users to fetch")
		}
		return nil, err
	}
	return user, nil
}

// GetUserByName implements repointer.RepoInter.
func (r *RepoImpl) GetByName(name string) (*models.User, error) {
	user := new(models.User)
	// err := utils.DB.NewSelect().Model(user).Where("name=?", name).Scan(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// return user, nil
	row := utils.DB.QueryRow("select * from users where name=?", name)
	if err := row.Scan(&user.UserID, &user.Name, &user.Height, &user.BirthDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no users to fetch")
		}
		return nil, err
	}
	return user, nil
}

func NewRepoImpl() *RepoImpl {
	return &RepoImpl{}
}
