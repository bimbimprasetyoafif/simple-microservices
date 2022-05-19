package comment

import (
	"context"
	"fmt"
	pb "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/comment"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"google.golang.org/grpc"
)

type Client struct {
	c pb.CommentServiceClient
}

func InitClient(url, port string) (provider.CommentClientProvider, error) {
	co, err := grpc.Dial(fmt.Sprintf("%s:%s", url, port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return &Client{}, err
	}

	return &Client{
		c: pb.NewCommentServiceClient(co),
	}, nil
}

func (o *Client) CreateComment(orgName string, value string) (*pb.CreateResp, error) {
	return o.c.Create(context.Background(), &pb.CreateReq{
		OrgSlug: orgName,
		Value:   value,
	})
}

func (o *Client) GetAllComment(orgName string) (*pb.GetAllResp, error) {
	return o.c.GetAll(context.Background(), &pb.GetAllReq{
		OrgSlug: orgName,
	})
}

func (o *Client) DeleteComment(orgName string) (*pb.DeleteResp, error) {
	return o.c.Delete(context.Background(), &pb.DeleteReq{
		OrgSlug: orgName,
	})
}
