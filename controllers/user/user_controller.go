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

func (u *Handler) QueryByID(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := u.SVC.QueryByID(id)
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
func (u *Handler) Delete(c *gin.Context) {
	strid := c.Param("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}
	res, err := u.SVC.Delete(id)
	errmessage := fmt.Sprintf("%s", err)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error deleting the user- handler",
			"data":    errmessage,
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully deleted the user",
		"data":    res,
	})
}
func (u *Handler) QueryByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	res, err := u.SVC.QueryByName(name)
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

func (u *Handler) QueryAll(c *gin.Context) {
	res, err := u.SVC.QueryAll()
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
func (u *Handler) Update(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := u.SVC.Update(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error updating the user- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully updated the user",
		"data":    res,
	})
}
func NewUserHandler(svc *user.SVCImpl) *Handler {
	return &Handler{SVC: svc}
}
