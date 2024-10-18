package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type(
	OrganizationUsecase interface{
		GetOrganizationList() ([]model.OrganizationResponse, error)
		CreateOrganization (organization model.Organization) (model.OrganizationResponse, error)
		UpdateOrganization (organization model.Organization, id uint) (model.OrganizationResponse, error)
		DeleteOrganization (organization model.Organization, id uint) error
	}

	organizationUsecase struct{
		or repository.OrganizationRepository
		ov validator.OrganizationValidator
	}
)

func NewOrganizationUsecase (or repository.OrganizationRepository, ov validator.OrganizationValidator) OrganizationUsecase {
	return &organizationUsecase{or,ov}
}

func (ou *organizationUsecase) GetOrganizationList () ([]model.OrganizationResponse, error) {

	organizations := []model.Organization{}
	if err := ou.or.GetOrganizationList(&organizations); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return nil, err
	}

	resOrganizations := []model.OrganizationResponse{}

	for _,v:= range(organizations){
		res := model.OrganizationResponse{
			ID: v.ID,
			Name: v.Name,
			AccessGiven: v.AccessGiven,
		}
		resOrganizations = append(resOrganizations, res)
	}

	return resOrganizations, nil
}

func (ou *organizationUsecase) CreateOrganization (organization model.Organization) (model.OrganizationResponse, error){

	if err := ou.ov.OrganizationValidate(organization); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrganization")
		return model.OrganizationResponse{}, err
	}

	if err := ou.or.CreateOrganization(&organization); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrganization")
		return model.OrganizationResponse{}, err
	}

	resOrganization := model.OrganizationResponse{
		ID: organization.ID,
		Name: organization.Name,
		AccessGiven: organization.AccessGiven,
	}

	return resOrganization, nil
}

func (ou *organizationUsecase) UpdateOrganization (organization model.Organization, id uint) (model.OrganizationResponse, error) {

	if err := ou.ov.OrganizationValidate(organization); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrganization")
		return model.OrganizationResponse{}, err
	}

	if err := ou.or.UpdateOrganization(&organization, id); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrganization")
		return model.OrganizationResponse{}, err
	}

	resOrganization := model.OrganizationResponse{
		ID: organization.ID,
		Name: organization.Name,
		AccessGiven: organization.AccessGiven,
	}
	return resOrganization, nil
}

func (ou *organizationUsecase) DeleteOrganization (organization model.Organization, id uint) error{

	if err := ou.or.DeleteOrganization(&organization, id); err!=nil{
		return err
	}

	return nil;
}



