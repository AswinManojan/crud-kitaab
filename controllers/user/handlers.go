package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/repositories/user/models"
	"github.com/sample-crud-app/services/user"
)

type Handler struct {
	SVC *user.SVCImpl
}

func (u *Handler) Create(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := u.SVC.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error creating the user- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully created the user",
		"data":    res,
	})
}

func (u *Handler) GetByID(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := u.SVC.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the user- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the user",
		"data":    res,
	})
}
func (u *Handler) GetByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	res, err := u.SVC.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the user- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the user",
		"data":    res,
	})
}

func (u *Handler) GetAll(c *gin.Context) {
	res, err := u.SVC.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the users",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the users",
		"data":    res,
	})
}

func NewUserHandler(svc *user.SVCImpl) *Handler {
	return &Handler{SVC: svc}
}
