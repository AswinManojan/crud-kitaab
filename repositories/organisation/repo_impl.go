package organisation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/utils"
)

type RepoImpl struct {
}

// CreateOrganization implements repointer.RepoInter.
func (r *RepoImpl) Create(organization *models.Organization) (*models.Organization, error) {
	// _, err := utils.DB.Exec("INSERT INTO Organization (legal_name,Alias,Country,Currency,Gstreg,gstin,state,pan) values ($1,$2,$3,$4,$5,$6,$7,$8)", organization.LegalName, organization.Alias, organization.Country, organization.Currency, organization.Gstreg, organization.Gstin, organization.State, organization.Pan)
	_, err := utils.DB.NewInsert().Model(organization).Exec(context.Background())
	if err != nil {
		fmt.Println("Error while creating organization, repo layer")
		return nil, err
	}
	return organization, nil
}

// DeleteOrganizaionByID implements repointer.RepoInter.
func (r *RepoImpl) DeleteByID(id int) (bool, error) {
	_, err := utils.DB.NewDelete().Model(&models.Organization{}).Where("id=?", id).Exec(context.Background())
	if err != nil {
		fmt.Println("Error while deleting organization, repo layer")
		return false, err
	}
	return true, nil
}

// GetOrganizationByID implements repointer.RepoInter.
func (r *RepoImpl) GetByID(id int) (*models.Organization, error) {
	organization := new(models.Organization)
	// err := utils.DB.NewSelect().Model(organisation).Where("id=?", id).Scan(context.Background())
	row := utils.DB.QueryRow("select * from organizations where id=?", id)
	if err := row.Scan(&organization.ID, &organization.LegalName, &organization.Alias, &organization.Country, &organization.Currency, &organization.Gstreg, &organization.Gstin, &organization.State, &organization.Pan, &organization.UserID, &organization.ParentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("organizationByID %d: no such organization", id)
		}
		return nil, err
	}
	return organization, nil
	// fmt.Println(id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// return organisation, nil
}

func (r *RepoImpl) GetAll() ([]models.Organization, error) {
	var organizations []models.Organization
	rows, _ := utils.DB.Query("select * from organizations")
	// fmt.Println(rows)
	for rows.Next() {
		var organization models.Organization
		if err := rows.Scan(&organization.ID, &organization.LegalName, &organization.Alias, &organization.Country, &organization.Currency, &organization.Gstreg, &organization.Gstin, &organization.State, &organization.Pan, &organization.UserID, &organization.ParentID); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no organisations to fetch")
			}
			return nil, err
		}
		organizations = append(organizations, organization)
	}
	// fmt.Println(rows)
	return organizations, nil
}

// GetOrganizationByName implements repointer.RepoInter.
func (r *RepoImpl) GetByName(name string) (*models.Organization, error) {
	var organization models.Organization
	row := utils.DB.QueryRow("select * from organizations where legal_name=?", name)
	if err := row.Scan(&organization.ID, &organization.LegalName, &organization.Alias, &organization.Country, &organization.Currency, &organization.Gstreg, &organization.Gstin, &organization.State, &organization.Pan, &organization.UserID, &organization.ParentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("organizationByName %s: no such organization", name)
		}
		return nil, err
	}
	return &organization, nil
	// organisation := new(models.Organization)
	// err := utils.DB.NewSelect().Model(organisation).Where("legal_name=?", name).Scan(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// return organisation, nil
}

// UpdateOrganizationByID implements repointer.RepoInter.
func (r *RepoImpl) Update(id int, organization *models.Organization) (*models.Organization, error) {
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
