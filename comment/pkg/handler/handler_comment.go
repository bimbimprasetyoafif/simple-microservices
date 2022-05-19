package handler

import (
	"context"
	pb "github.com/bimbimprasetyoafif/comment/pkg/proto"
	"google.golang.org/grpc"
)

type commentHandler struct {
	uc UsecaseProvider
}

func (c *commentHandler) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error) {
	err := c.uc.DeleteComment(req.GetOrgSlug())
	if err != nil {
		return &pb.DeleteResp{
			Status:  500,
			Message: "something wrong",
		}, nil
	}

	return &pb.DeleteResp{
		Status:  200,
		Message: "success",
	}, nil
}

func (c *commentHandler) GetAll(ctx context.Context, req *pb.GetAllReq) (*pb.GetAllResp, error) {
	res, err := c.uc.GetAllComment(req.GetOrgSlug())
	if err != nil {
		return &pb.GetAllResp{
			Status:  500,
			Message: "something wrong",
			Data:    nil,
		}, nil
	}

	if res == nil {
		return &pb.GetAllResp{
			Status:  int32(404),
			Message: "org not found",
			Data:    nil,
		}, nil
	}

	data := make([]string, 0)
	for _, v := range res {
		data = append(data, v.Value)
	}
	return &pb.GetAllResp{
		Status:  int32(200),
		Message: "success",
		Data:    data,
	}, nil
}

func (c *commentHandler) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	success, err := c.uc.CreateComment(req.GetOrgSlug(), req.GetValue())
	if err != nil {
		return &pb.CreateResp{
			Status:  500,
			Message: "something wrong",
		}, nil
	}

	if !success {
		return &pb.CreateResp{
			Status:  404,
			Message: "org not found",
		}, nil
	}

	return &pb.CreateResp{
		Status:  200,
		Message: "success",
	}, nil
}

func NewGrpcHandler(server *grpc.Server, provider UsecaseProvider) {
	pb.RegisterCommentServiceServer(server, &commentHandler{
		uc: provider,
	})
}
