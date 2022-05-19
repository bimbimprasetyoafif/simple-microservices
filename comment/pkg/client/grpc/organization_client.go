package grpc

import (
	"context"
	"fmt"
	pb "github.com/bimbimprasetyoafif/comment/pkg/proto/organization"
	"github.com/bimbimprasetyoafif/comment/pkg/usecase"
	"google.golang.org/grpc"
)

type OrganizationClient struct {
	c pb.OrganizationServiceClient
}

func InitOrgClient(url, port string) (usecase.OrganizationClientProvider, error) {
	co, err := grpc.Dial(fmt.Sprintf("%s:%s", url, port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return &OrganizationClient{}, err
	}

	return &OrganizationClient{
		c: pb.NewOrganizationServiceClient(co),
	}, nil
}

func (o *OrganizationClient) CheckOrg(orgName string) (*pb.CheckResp, error) {
	return o.c.CheckAvaibility(context.Background(), &pb.CheckReq{
		Slug: orgName,
	})
}
