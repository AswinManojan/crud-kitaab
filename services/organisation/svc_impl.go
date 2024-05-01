package organisation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	repo "github.com/sample-crud-app/repositories/organisation"
	"github.com/sample-crud-app/repositories/organisation/models"
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

type SVCImpl struct {
	Repo *repo.RepoImpl
}

// CreateOrganization implements svcinter.SVCInter.
func (s *SVCImpl) Create(organization *models.Organization) (*models.Organization, error) {
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

	org, err := s.Repo.Create(organization)
	if err != nil {
		fmt.Println("Error creating organization in svc layer")
		return nil, err
	}
	return org, nil
}

// DeleteOrganizaionByID implements svcinter.SVCInter.
func (s *SVCImpl) DeleteByID(id int) (bool, error) {
	result, err := s.Repo.DeleteByID(id)
	if err != nil {
		fmt.Println("Error deleting organization in svc layer")
		return result, err
	}
	return result, nil
}

// GetOrganizationByID implements svcinter.SVCInter.
func (s *SVCImpl) GetByID(id int) (*models.Organization, error) {
	org, err := s.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Error finding organization in svc layer")
		fmt.Println(err)
		return nil, err
	}
	return org, nil
}

// GetOrganizationByName implements svcinter.SVCInter.
func (s *SVCImpl) GetByName(name string) (*models.Organization, error) {
	org, err := s.Repo.GetByName(name)
	if err != nil {
		fmt.Println("Error finding organization in svc layer")
		return nil, err
	}
	return org, nil
}

// UpdateOrganization implements svcinter.SVCInter.
func (s *SVCImpl) Update(id int, organization *models.Organization) (*models.Organization, error) {
	organization.ID = uint(id)
	org, err := s.Repo.Update(id, organization)
	if err != nil {
		fmt.Println("Error updating organization in svc layer")
		return nil, err
	}
	return org, nil
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

func NewSVCImpl(repo *repo.RepoImpl) *SVCImpl {
	return &SVCImpl{Repo: repo}
}
