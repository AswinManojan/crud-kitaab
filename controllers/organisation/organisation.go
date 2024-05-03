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


// CreateOrganisations	godoc
// @Summary				Create Organisations
// @Description			Saves Organisations data in DB
// @Produce				application/json
// @Tags				Organisation
// @Param				Organisation-Details	body models.Organization true "Organisation Data"
// @Success				202 {object} models.Organization
// @Failure				400 {object} string "Error Details"
// @Router 				/organizations [post]
func (o *OrganisationController) Create(c *gin.Context) {
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := organisationService.Create(orgn)
	user.ResponseMessage(c, res, err)

}

// QueryOrganisationByID	godoc
// @Summary				Query Organisations By ID
// @Description			Get Organisations details by ID
// @Produce				application/json
// @Tags				Organisation
// @Param				Organisation-ID	path int true "Organisation ID"
// @Success				202 {object} models.Organization
// @Failure				400 {object} string "Error Details"
// @Router 				/organizations/{Organisation-ID} [get]
func (o *OrganisationController) QueryByID(c *gin.Context) {
	strid := c.Param("Organisation-ID")
	id, err := strconv.Atoi(strid)
	if err != nil {
		user.ResponseMessage(c, nil, err)
	}
	res, err := organisationService.QueryByID(id)
	user.ResponseMessage(c, res, err)
}

// QueryOrganisationByName	godoc
// @Summary				Query Organisations By Name
// @Description			Get Organisations details by Name
// @Produce				application/json
// @Tags				Organisation
// @Param				Organisation-Name query string true "Organisation-Name"
// @Success				202 {object} models.Organization
// @Failure				400 {object} string "Error Details"
// @Router 				/organization [get]
func (o *OrganisationController) QueryByName(c *gin.Context) {
	name := c.DefaultQuery("Organisation-Name", "")
	res, err := organisationService.QueryByName(name)
	user.ResponseMessage(c, res, err)
}

// QueryAllOrganisations	godoc
// @Summary				Query All Organisations
// @Description			Get All Organisations details
// @Produce				application/json
// @Tags				Organisation
// @Success				202 {object} []models.Organization
// @Failure				400 {object} string "Error Details"
// @Router 				/organizations [get]
func (o *OrganisationController) QueryAll(c *gin.Context) {
	res, err := organisationService.QueryAll()
	user.ResponseMessage(c, res, err)
}

// UpdateOrganisationByID	godoc
// @Summary				Update Organisation By ID
// @Description			Update Organisations details by ID
// @Produce				application/json
// @Tags				Organisation
// @Param				Organisation-ID	path int true "Organisation ID"
// @Param				Organisation-Details	body models.Organization true "Organisation Data"
// @Success				202 {object} models.Organization
// @Failure				400 {object} string "Error Details"
// @Router 				/organizations/{Organisation-ID} [put]
func (o *OrganisationController) Update(c *gin.Context) {
	strid := c.Param("Organisation-ID")
	id, _ := strconv.Atoi(strid)
	var orgn *models.Organization
	if err := c.BindJSON(&orgn); err != nil {
		log.Println("Error binding the JSON data")
		return
	}
	res, err := organisationService.Update(id, orgn)
	user.ResponseMessage(c, res, err)
}

// DeleteOrganisationByID	godoc
// @Summary				Delete Organisation By ID
// @Description			Delete Organisations details by ID
// @Produce				application/json
// @Tags				Organisation
// @Param				Organisation-ID	path int true "Organisation ID"
// @Success				202 {object} bool
// @Failure				400 {object} string "Error Details"
// @Router 				/organizations/{Organisation-ID} [delete]
func (o *OrganisationController) Delete(c *gin.Context) {
	strid := c.Param("Organisation-ID")
	id, _ := strconv.Atoi(strid)
	res, err := organisationService.Delete(id)
	user.ResponseMessage(c, res, err)
}
