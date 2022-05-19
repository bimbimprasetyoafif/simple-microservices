package repository

import (
	"github.com/bimbimprasetyoafif/organization/pkg/model"
	"gorm.io/gorm"
)

type orgPostgres struct {
	db *gorm.DB
}

func (o *orgPostgres) RegisterOrg(orgName, orgSlug string) (model.Organization, error) {
	var mOrg model.Organization
	mOrg.Name = orgName
	mOrg.Slug = orgSlug
	err := o.db.Create(&mOrg).Error

	return mOrg, err
}

func (o *orgPostgres) CheckOrgName(orgNameSlug string) *model.Organization {
	var mOrg model.Organization

	o.db.Where("slug = ?", orgNameSlug).First(&mOrg)

	return &mOrg
}

func NewOrgMysql(db *gorm.DB) *orgPostgres {
	return &orgPostgres{
		db: db,
	}
}
