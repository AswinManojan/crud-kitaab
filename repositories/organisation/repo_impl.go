package organisation

import (
	"context"
	"fmt"

	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/utils"
)

type RepoImpl struct {
}

// CreateOrganization implements repointer.RepoInter.
func (r *RepoImpl) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	// _, err := utils.DB.Exec("INSERT INTO Organization (legal_name,Alias,Country,Currency,Gstreg,gstin,state,pan) values ($1,$2,$3,$4,$5,$6,$7,$8)", organization.LegalName, organization.Alias, organization.Country, organization.Currency, organization.Gstreg, organization.Gstin, organization.State, organization.Pan)
	_, err := utils.DB.NewInsert().Model(organization).Exec(context.Background())
	if err != nil {
		fmt.Println("Error while creating organization, repo layer")
		return nil, err
	}
	return organization, nil
}

// DeleteOrganizaionByID implements repointer.RepoInter.
func (r *RepoImpl) DeleteOrganizaionByID(id int) (bool, error) {
	_, err := utils.DB.NewDelete().Model(&models.Organization{}).Where("id=?", id).Exec(context.Background())
	if err != nil {
		fmt.Println("Error while deleting organization, repo layer")
		return false, err
	}
	return true, nil
}

// GetOrganizationByID implements repointer.RepoInter.
func (r *RepoImpl) GetOrganizationByID(id int) (*models.Organization, error) {
	organisation := new(models.Organization)
	err := utils.DB.NewSelect().Model(organisation).Where("id=?", id).Scan(context.Background())
	// row := utils.DB.QueryRow("select * from organization where id=$1", id)
	// fmt.Println(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return organisation, nil
}

// GetOrganizationByName implements repointer.RepoInter.
func (r *RepoImpl) GetOrganizationByName(name string) (*models.Organization, error) {
	// var organization models.Organization
	// row := utils.DB.QueryRow("select * from organization where legal_name=$1", name)
	// if err := row.Scan(&organization.LegalName, &organization.Alias, &organization.Country, &organization.Currency, &organization.Gstreg, &organization.Gstin, &organization.State, &organization.Pan, &organization.ID); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, fmt.Errorf("organizationByName %s: no such organization", name)
	// 	}
	// 	return nil, err
	// }
	// return &organization, nil
	organisation := new(models.Organization)
	err := utils.DB.NewSelect().Model(organisation).Where("legal_name=?", name).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return organisation, nil
}

// UpdateOrganizationByID implements repointer.RepoInter.
func (r *RepoImpl) UpdateOrganization(id int, organization *models.Organization) (*models.Organization, error) {
	// fmt.Println(organization)
	_, err := utils.DB.NewUpdate().Model(organization).Set("legal_name=?", organization.LegalName).Set("alias=?", organization.Alias).
		Set("country=?", organization.Country).
		Set("currency=?", organization.Currency).
		Set("gstin=?", organization.Gstin).
		Set("gstreg=?", organization.Gstreg).
		Set("state=?", organization.State).
		Set("pan=?", organization.Pan).
		Set("id=?", id).
		Where("id=?", id).
		Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return organization, nil
}

func NewRepoImpl() *RepoImpl {
	return &RepoImpl{}
}
