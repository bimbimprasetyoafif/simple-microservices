package grpc

import (
	"context"
	"github.com/bimbimprasetyoafif/organization/pkg/handler"
	pb "github.com/bimbimprasetyoafif/organization/pkg/proto"
	"google.golang.org/grpc"
)

type handlerOrg struct {
	usecase handler.UsecaseProvider
}

func (h *handlerOrg) CheckAvaibility(ctx context.Context, req *pb.CheckReq) (*pb.CheckResp, error) {
	res := h.usecase.CheckOrganization(ctx, req.GetSlug())
	return &pb.CheckResp{
		OrgId:   int32(res.ID),
		IsExist: !(res.ID <= 0),
	}, nil
}

func (h *handlerOrg) Create(ctx context.Context, req *pb.CreateReq) (*pb.MessageResp, error) {
	res, err := h.usecase.CreateOrganization(ctx, req.GetName())
	if err != nil {

		// lil bit tricky to handle sql state postgres duplicate key
		return &pb.MessageResp{
			Status:  400,
			Message: "org already exist",
			Data:    nil,
		}, nil
	}

	return &pb.MessageResp{
		Status:  200,
		Message: "success",
		Data: &pb.Organization{
			Name: res.Name,
			Slug: res.Slug,
		},
	}, nil
}

func NewGrpcHandler(server *grpc.Server, usecase handler.UsecaseProvider) {
	pb.RegisterOrganizationServiceServer(server, &handlerOrg{
		usecase: usecase,
	})
}
