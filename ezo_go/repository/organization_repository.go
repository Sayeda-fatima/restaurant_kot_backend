package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type OrganizationRepository interface{
	GetOrganizationList (organization *[]model.Organization, page int) error
	CreateOrganization (organization *model.Organization) error
	UpdateOrganization (organization *model.Organization, id uint) error
	DeleteOrganization (organization *model.Organization, id uint) error
}

type organizationRepository struct{
	db *gorm.DB
}

func NewOrganizationRepository (db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (or *organizationRepository) GetOrganizationList(organization *[]model.Organization, page int) error{

	if err := or.db.Scopes(common.Paginate(page)).Find(organization).Error; err!=nil{
		return err
	}

	return nil
}

func (or *organizationRepository) CreateOrganization (organization *model.Organization) error {

	if err := or.db.Create(organization).Error; err!=nil{
		return err
	}
	return nil
}

func (or *organizationRepository) UpdateOrganization (organization *model.Organization, id uint) error {

	result := or.db.Model(organization).Where("id=?",id).Updates(organization)

	if err := result.Error; err!=nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}

func (or *organizationRepository) DeleteOrganization (organization *model.Organization, id uint) error {

	result:=or.db.Model(organization).Where("id=?",id).Update("is_deleted", 1)

	if err := result.Error; err!=nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("record does not exist")
	}
	
	return nil
}