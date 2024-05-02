package organisation

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sample-crud-app/controllers/user"
	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/services/organisation"
)

var organisationService *organisation.OrganisationService

type OrganisationController struct{}

func (o *OrganisationController) Create(c *gin.Context) {
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := organisationService.Create(orgn)
	user.ResponseMessage(c, res, err)

}

func (o *OrganisationController) Update(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := organisationService.Update(id, orgn)
	user.ResponseMessage(c, res, err)
}

func (o *OrganisationController) QueryByID(c *gin.Context) {
	strid := c.Param("id")
	// fmt.Println(strid)
	id, _ := strconv.Atoi(strid)
	// fmt.Println(id)
	res, err := organisationService.QueryByID(id)
	user.ResponseMessage(c, res, err)
}

func (o *OrganisationController) QueryByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	res, err := organisationService.QueryByName(name)
	user.ResponseMessage(c, res, err)
}

func (o *OrganisationController) QueryAll(c *gin.Context) {
	res, err := organisationService.QueryAll()
	user.ResponseMessage(c, res, err)
}

func (o *OrganisationController) Delete(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)
	res, err := organisationService.Delete(id)
	user.ResponseMessage(c, res, err)
}
