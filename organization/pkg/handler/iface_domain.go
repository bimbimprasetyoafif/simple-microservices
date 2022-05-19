package handler

import (
	"context"
	"github.com/bimbimprasetyoafif/organization/pkg/model"
)

type UsecaseProvider interface {
	CreateOrganization(ctx context.Context, orgName string) (model.Organization, error)
	CheckOrganization(ctx context.Context, orgSlug string) *model.Organization
}
