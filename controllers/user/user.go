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

func (u *UserController) Create(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := userService.Create(user)
	ResponseMessage(c, res, err)
}

func (u *UserController) QueryByID(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := userService.QueryByID(id)
	ResponseMessage(c, res, err)
}
func (u *UserController) Delete(c *gin.Context) {
	strid := c.Param("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}
	res, err := userService.Delete(id)
	ResponseMessage(c, res, err)
}

func (u *UserController) QueryAll(c *gin.Context) {
	res, err := userService.QueryAll()
	ResponseMessage(c, res, err)
}
func (u *UserController) Update(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := userService.Update(id, user)
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
