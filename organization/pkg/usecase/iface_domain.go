package usecase

import (
	"github.com/bimbimprasetyoafif/organization/pkg/model"
)

type RepositoryProvider interface {
	RegisterOrg(orgName, orgSLug string) (model.Organization, error)
	CheckOrgName(orgNameSlug string) *model.Organization
}
