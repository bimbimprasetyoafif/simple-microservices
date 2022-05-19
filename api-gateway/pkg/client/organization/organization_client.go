package organization

import (
	"context"
	"fmt"
	pb "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/organization"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"google.golang.org/grpc"
)

type Client struct {
	c pb.OrganizationServiceClient
}

func InitClient(url, port string) (provider.OrgClientProvider, error) {
	co, err := grpc.Dial(fmt.Sprintf("%s:%s", url, port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return &Client{}, err
	}

	return &Client{
		c: pb.NewOrganizationServiceClient(co),
	}, nil
}

func (o *Client) RegisterOrganization(orgName string) (*pb.MessageResp, error) {
	return o.c.Create(context.Background(), &pb.CreateReq{
		Name: orgName,
	})
}
