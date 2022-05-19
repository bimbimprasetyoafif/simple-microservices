package testing

import pb "github.com/bimbimprasetyoafif/comment/pkg/proto/organization"

type ClientOrgMock struct {
}

func (o *ClientOrgMock) CheckOrg(orgName string) (*pb.CheckResp, error) {
	return &pb.CheckResp{
		IsExist: true,
		OrgId:   1,
	}, nil
}
