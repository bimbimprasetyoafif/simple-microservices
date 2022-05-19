package provider

import (
	pbCom "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/comment"
	pbOrg "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/organization"
)

type ClientCommentMock struct {
	CreateFunc func(orgName string, value string) (*pbCom.CreateResp, error)
	GetFunc    func(orgName string) (*pbCom.GetAllResp, error)
	DelFunc    func(orgName string) (*pbCom.DeleteResp, error)
}

func (c *ClientCommentMock) CreateComment(orgName string, value string) (*pbCom.CreateResp, error) {
	return c.CreateFunc(orgName, value)
}
func (c *ClientCommentMock) GetAllComment(orgName string) (*pbCom.GetAllResp, error) {
	return c.GetFunc(orgName)
}
func (c *ClientCommentMock) DeleteComment(orgName string) (*pbCom.DeleteResp, error) {
	return c.DelFunc(orgName)
}

type ClientOrgMock struct {
	RegisFunc func(orgName string) (*pbOrg.MessageResp, error)
}

func (c *ClientOrgMock) RegisterOrganization(orgName string) (*pbOrg.MessageResp, error) {
	return c.RegisFunc(orgName)
}
