package handler

import (
	"context"
	"github.com/bimbimprasetyoafif/organization/pkg/model"
)

type UsecaseMock struct {
	CreateF func(ctx context.Context, orgName string) (model.Organization, error)
	CheckF  func(ctx context.Context, orgSlug string) *model.Organization
}

func (u *UsecaseMock) CreateOrganization(ctx context.Context, orgName string) (model.Organization, error) {
	return u.CreateF(ctx, orgName)
}
func (u *UsecaseMock) CheckOrganization(ctx context.Context, orgSlug string) *model.Organization {
	return u.CheckF(ctx, orgSlug)
}
