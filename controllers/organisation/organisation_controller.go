package organisation

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/services/organisation"
)

type Handler struct {
	SVC *organisation.SVCImpl
}

func (o *Handler) Create(c *gin.Context) {
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := o.SVC.Create(orgn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error creating the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully created the Organization",
		"data":    res,
	})
}

func (o *Handler) Update(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := o.SVC.Update(id, orgn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error updating the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully updated the Organization",
		"data":    res,
	})
}
func (o *Handler) QueryByID(c *gin.Context) {
	strid := c.Param("id")
	// fmt.Println(strid)
	id, _ := strconv.Atoi(strid)
	// fmt.Println(id)
	res, err := o.SVC.QueryByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the Organization",
		"data":    res,
	})
}
func (o *Handler) QueryByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	res, err := o.SVC.QueryByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the Organization",
		"data":    res,
	})
}
func (o *Handler) QueryAll(c *gin.Context) {
	res, err := o.SVC.QueryAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the Organization",
		"data":    res,
	})
}
func (o *Handler) Delete(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := o.SVC.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Error finding the organization- handler",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully found the Organization",
		"data":    res,
	})
}
func NewOrganizationHandler(svc *organisation.SVCImpl) *Handler {
	return &Handler{SVC: svc}
}