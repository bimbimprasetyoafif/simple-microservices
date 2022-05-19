package usecase

import (
	"github.com/bimbimprasetyoafif/comment/pkg/model"
	pb "github.com/bimbimprasetyoafif/comment/pkg/proto/organization"
)

type CommentProvider interface {
	InsertOne(orgId int, commentValue string) error
	GetAllByOrg(orgID int) ([]model.Comment, error)
	DeleteByOrg(orgID int) error
}

type OrganizationClientProvider interface {
	CheckOrg(orgName string) (*pb.CheckResp, error)
}
