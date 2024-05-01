package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/services/organisation"
)

type OrganizationHandler struct {
	SVC *services.SVCImpl
}

func (o *OrganizationHandler) CreateOrganizationHandler(c *gin.Context) {
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := o.SVC.CreateOrganization(orgn)
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

func (o *OrganizationHandler) UpdateOrganizationHandler(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := o.SVC.UpdateOrganization(id, orgn)
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
func (o *OrganizationHandler) GetOrganizationByIDHandler(c *gin.Context) {
	strid := c.Param("id")
	// fmt.Println(strid)
	id, _ := strconv.Atoi(strid)
	// fmt.Println(id)
	res, err := o.SVC.GetOrganizationByID(id)
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
func (o *OrganizationHandler) GetOrganizationByNameHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	// fmt.Println(name)
	res, err := o.SVC.GetOrganizationByName(name)
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
func (o *OrganizationHandler) DeleteOrganizaionByIDHandler(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := o.SVC.DeleteOrganizaionByID(id)
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
func NewOrganizationHandler(svc *services.SVCImpl) *OrganizationHandler {
	return &OrganizationHandler{SVC: svc}
}
