package organisation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sample-crud-app/repositories/organisation"
	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/repositories/user"
)

type Address struct {
	StateName string `json:"stateName"`
}

type Data struct {
	GSTIN     string    `json:"gstin"`
	LegalName string    `json:"legalName"`
	Addresses []Address `json:"addresses"`
}

type Response struct {
	Data    Data `json:"data"`
	Success bool `json:"success"`
}

var repository *organisation.OrganisationRepository
var userRepository *user.UserRepository

type OrganisationService struct{}

// CreateOrganization implements svcinter.SVCInter.
func (s *OrganisationService) Create(organization *models.Organization) (*models.Organization, error) {
	if _, err := repository.QueryByID(int(organization.ParentID)); organization.ParentID > 0 && err != nil {
		return nil, errors.New("no existing organisation for as the provided parent organisation")
	}
	_, err := userRepository.QueryByID(int(organization.UserID))
	if err != nil {
		return nil, errors.New("no existing user for adding the organisation")
	}
	status := GSTVerifier(organization.Gstin)
	if status != "200 OK" {
		errmessage := fmt.Sprintf("Error - Invalid GST Number - %s", status)
		return nil, errors.New(errmessage)
	}
	if organization.Pan == "" {
		organization.Pan = organization.Gstin[2:12]
	}
	if organization.Gstin[2:12] != organization.Pan {
		return nil, errors.New("error - Invalid PAN Number provided")
	}
	response := GSTInfoFetch(organization.Gstin)
	if organization.LegalName == "" {
		organization.LegalName = response.Data.LegalName
	}
	if organization.LegalName != response.Data.LegalName {
		return nil, errors.New("error - Invalid Legal Name provided")
	}
	if organization.State == "" {
		organization.State = response.Data.Addresses[0].StateName
	}
	if organization.State != response.Data.Addresses[0].StateName {
		return nil, errors.New("error - Invalid State Name provided")
	}

	org, err := repository.Create(organization)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// DeleteOrganizaionByID implements svcinter.SVCInter.
func (s *OrganisationService) Delete(id int) (bool, error) {
	return repository.Delete(id)
}

// GetOrganizationByID implements svcinter.SVCInter.
func (s *OrganisationService) QueryByID(id int) (*models.Organization, error) {
	return repository.QueryByID(id)
}

// GetOrganizationByName implements svcinter.SVCInter.
func (s *OrganisationService) QueryByName(name string) (*models.Organization, error) {
	return repository.QueryByName(name)
}

func (s *OrganisationService) QueryAll() ([]models.Organization, error) {
	return repository.QueryAll()
}

// UpdateOrganization implements svcinter.SVCInter.
func (s *OrganisationService) Update(id int, organization *models.Organization) (*models.Organization, error) {
	organization.ID = uint(id)
	return repository.Update(id, organization)
}

func GSTVerifier(gst string) string {
	str := fmt.Sprintf("https://devapi.finsights.biz/finsightsgstinapi/v1/%s/details", gst)
	resp, err := http.Get(str)
	if err != nil {
		fmt.Println("Error while verifying GST")
	}
	// body, err := io.ReadAll(resp.Body)
	return resp.Status
}

func GSTInfoFetch(gst string) Response {
	var response Response
	str := fmt.Sprintf("https://devapi.finsights.biz/finsightsgstinapi/v1/%s/details", gst)
	resp, err := http.Get(str)
	if err != nil {
		fmt.Println("Error while fetching GST info")
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)
	return response
}
