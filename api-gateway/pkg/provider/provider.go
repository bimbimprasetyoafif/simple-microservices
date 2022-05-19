package provider

import (
	pbCom "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/comment"
	pbOrg "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/organization"
)

type OrgClientProvider interface {
	RegisterOrganization(orgName string) (*pbOrg.MessageResp, error)
}

type CommentClientProvider interface {
	CreateComment(orgName string, value string) (*pbCom.CreateResp, error)
	GetAllComment(orgName string) (*pbCom.GetAllResp, error)
	DeleteComment(orgName string) (*pbCom.DeleteResp, error)
}
