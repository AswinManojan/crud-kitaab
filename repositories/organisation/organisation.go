package organisation

import (
	"context"

	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/utils"
)

type OrganisationRepository struct{}

// CreateOrganization implements repointer.RepoInter.
func (r *OrganisationRepository) Create(organization *models.Organization) (*models.Organization, error) {
	_, err := utils.DB.NewInsert().Model(organization).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return organization, nil
}

// DeleteOrganizaionByID implements repointer.RepoInter.
func (r *OrganisationRepository) Delete(id int) (bool, error) {
	_, err := utils.DB.NewDelete().Model(&models.Organization{}).Where("organisation_id=?", id).Exec(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetOrganizationByID implements repointer.RepoInter.
func (r *OrganisationRepository) QueryByID(id int) (*models.Organization, error) {
	organization := new(models.Organization)
	err := utils.DB.NewRaw("select * from organizations where organisation_id=?", id).Scan(context.Background(), organization)
	return organization, err
}

func (r *OrganisationRepository) QueryAll() ([]models.Organization, error) {
	var organizations []models.Organization
	err := utils.DB.NewRaw("select * from organizations").Scan(context.Background(), &organizations)
	return organizations, err
}

// GetOrganizationByName implements repointer.RepoInter.
func (r *OrganisationRepository) QueryByName(name string) (*models.Organization, error) {
	var organization models.Organization
	err := utils.DB.NewRaw("select * from organizations where legal_name=?", name).Scan(context.Background(), &organization)
	return &organization, err
}

// UpdateOrganizationByID implements repointer.RepoInter.
func (r *OrganisationRepository) Update(id int, organization *models.Organization) (*models.Organization, error) {
	// fmt.Println(organization)
	_, err := utils.DB.NewUpdate().Model(organization).Set("legal_name=?", organization.LegalName).Set("alias=?", organization.Alias).
		Set("country=?", organization.Country).
		Set("currency=?", organization.Currency).
		Set("gstin=?", organization.Gstin).
		Set("gstreg=?", organization.Gstreg).
		Set("state=?", organization.State).
		Set("pan=?", organization.Pan).
		Set("organisation_id=?", id).
		Where("organisation_id=?", id).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return organization, nil
}
