package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/bimbimprasetyoafif/organization/pkg/handler"
	"github.com/bimbimprasetyoafif/organization/pkg/model"
	pb "github.com/bimbimprasetyoafif/organization/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/gorm"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateOrg(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	mockUsecase := handler.UsecaseMock{}
	pb.RegisterOrganizationServiceServer(s, &handlerOrg{
		&mockUsecase,
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println("error mock dialer : ", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer func() {
		conn.Close()
		s.GracefulStop()
	}()

	testTable := []struct {
		name            string
		f               func(ctx context.Context, orgName string) (model.Organization, error)
		expectedStatus  int
		expectedMessage string
	}{
		{
			name: "org already exist",
			f: func(ctx context.Context, orgName string) (model.Organization, error) {
				return model.Organization{}, errors.New("error")
			},
			expectedStatus:  400,
			expectedMessage: "org already exist",
		},
		{
			name: "success",
			f: func(ctx context.Context, orgName string) (model.Organization, error) {
				return model.Organization{
					Name: "abc",
					Slug: "def",
				}, nil
			},
			expectedStatus:  200,
			expectedMessage: "success",
		},
	}

	cl := pb.NewOrganizationServiceClient(conn)
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			mockUsecase.CreateF = v.f
			resp, err := cl.Create(context.Background(), &pb.CreateReq{})
			assert.NoError(t, err)
			assert.Equal(t, v.expectedStatus, int(resp.GetStatus()))
		})
	}
}

func TestCheckAvailability(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	mockUsecase := handler.UsecaseMock{}
	pb.RegisterOrganizationServiceServer(s, &handlerOrg{
		&mockUsecase,
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println("error mock dialer : ", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer func() {
		conn.Close()
		s.GracefulStop()
	}()

	testTable := []struct {
		name           string
		f              func(ctx context.Context, orgSlug string) *model.Organization
		expectedStatus bool
		expectedID     int
	}{
		{
			name: "org exist",
			f: func(ctx context.Context, orgSlug string) *model.Organization {
				return &model.Organization{
					Model: gorm.Model{
						ID: 1,
					},
					Name: "abc",
					Slug: "def",
				}
			},
			expectedStatus: true,
			expectedID:     1,
		},
		{
			name: "org not exist",
			f: func(ctx context.Context, orgSlug string) *model.Organization {
				return &model.Organization{
					Model: gorm.Model{
						ID: 0,
					},
					Name: "",
					Slug: "",
				}
			},
			expectedStatus: false,
			expectedID:     0,
		},
	}

	cl := pb.NewOrganizationServiceClient(conn)
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			mockUsecase.CheckF = v.f
			resp, err := cl.CheckAvaibility(context.Background(), &pb.CheckReq{})
			assert.NoError(t, err)
			assert.Equal(t, v.expectedStatus, resp.GetIsExist())
			assert.Equal(t, v.expectedID, int(resp.GetOrgId()))
		})
	}
}
