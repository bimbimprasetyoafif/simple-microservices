package usecase

import (
	"context"
	"github.com/bimbimprasetyoafif/organization/pkg/model"
	"github.com/gosimple/slug"
)

type orgUsecase struct {
	repo RepositoryProvider
}

func (o *orgUsecase) CreateOrganization(ctx context.Context, orgName string) (model.Organization, error) {
	// Let's assume this line bellow for generate identity of org such a slug
	orgSlug := slug.Make(orgName)

	return o.repo.RegisterOrg(orgName, orgSlug)
}

func (o *orgUsecase) CheckOrganization(ctx context.Context, orgSlug string) *model.Organization {
	return o.repo.CheckOrgName(orgSlug)
}

func NewOrgUsecase(repo RepositoryProvider) *orgUsecase {
	return &orgUsecase{
		repo: repo,
	}
}
