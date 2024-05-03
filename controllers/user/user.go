package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/repositories/user/models"
	"github.com/sample-crud-app/services/user"
)

type UserController struct{}

var userService *user.UserService

// CreateUser	godoc
// @Summary				Create User
// @Description			Saves User data in DB
// @Produce				application/json
// @Tags				Users
// @Param				User-Details	body models.User true "User Data"
// @Success				202 {object} models.User
// @Failure				400 {object} string "Error Details"
// @Router 				/users [post]
func (u *UserController) Create(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := userService.Create(user)
	ResponseMessage(c, res, err)
}

// QueryUserByID	godoc
// @Summary				Query Users By ID
// @Description			Get Users details by ID
// @Produce				application/json
// @Tags				Users
// @Param				user-ID	path int true "Users ID"
// @Success				202 {object} models.User
// @Failure				400 {object} string "Error Details"
// @Router 				/users/{user-ID} [get]
func (u *UserController) QueryByID(c *gin.Context) {
	strid := c.Param("user-ID")
	id, _ := strconv.Atoi(strid)
	res, err := userService.QueryByID(id)
	ResponseMessage(c, res, err)
}

// QueryAllUsers	godoc
// @Summary				Query All Users
// @Description			Get All Users details
// @Produce				application/json
// @Tags				Users
// @Success				202 {object} []models.User
// @Failure				400 {object} string "Error Details"
// @Router 				/users [get]
func (u *UserController) QueryAll(c *gin.Context) {
	res, err := userService.QueryAll()
	ResponseMessage(c, res, err)
}

// UpdateUserByID	godoc
// @Summary				Update User By ID
// @Description			Update User details by ID
// @Produce				application/json
// @Tags				Users
// @Param				user-ID	path int true "Users ID"
// @Param				User-Details	body models.User true "User Data"
// @Success				202 {object} models.User
// @Failure				400 {object} string "Error Details"
// @Router 				/users/{user-ID} [put]
func (u *UserController) Update(c *gin.Context) {
	strid := c.Param("user-ID")
	id, _ := strconv.Atoi(strid)
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := userService.Update(id, user)
	ResponseMessage(c, res, err)
}

// DeleteUserByID	godoc
// @Summary				Delete User By ID
// @Description			Delete Users details by ID
// @Produce				application/json
// @Tags				Users
// @Param				user-ID	path int true "User ID"
// @Success				202 {object} bool
// @Failure				400 {object} string "Error Details"
// @Router 				/users/{user-ID} [delete]
func (u *UserController) Delete(c *gin.Context) {
	strid := c.Param("user-ID")
	id, err := strconv.Atoi(strid)
	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}
	res, err := userService.Delete(id)
	ResponseMessage(c, res, err)
}
func ResponseMessage(ctx *gin.Context, res any, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   res,
	})
}
