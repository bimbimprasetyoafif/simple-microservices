package handler

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/bimbimprasetyoafif/comment/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreate(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	mockUsecase := UsecaseMock{}
	pb.RegisterCommentServiceServer(s, &commentHandler{
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
		f               func(orgSlug, commentValue string) (bool, error)
		expectedStatus  int
		expectedMessage string
	}{
		{
			name: "org already exist",
			f: func(orgSlug, commentValue string) (bool, error) {
				return false, nil
			},
			expectedStatus:  404,
			expectedMessage: "org already exist",
		},
		{
			name: "success",
			f: func(orgSlug, commentValue string) (bool, error) {
				return true, nil
			},
			expectedStatus:  200,
			expectedMessage: "success",
		},
		{
			name: "internal error",
			f: func(orgSlug, commentValue string) (bool, error) {
				return false, errors.New("error")
			},
			expectedStatus:  500,
			expectedMessage: "something wrong",
		},
	}

	cl := pb.NewCommentServiceClient(conn)
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			mockUsecase.CreateFunc = v.f
			resp, err := cl.Create(context.Background(), &pb.CreateReq{})
			assert.NoError(t, err)
			assert.Equal(t, v.expectedStatus, int(resp.GetStatus()))
		})
	}
}

func TestDelete(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	mockUsecase := UsecaseMock{}
	pb.RegisterCommentServiceServer(s, &commentHandler{
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
		f               func(orgSlug string) error
		expectedStatus  int
		expectedMessage string
	}{
		{
			name: "success",
			f: func(orgSlug string) error {
				return nil
			},
			expectedStatus:  200,
			expectedMessage: "success",
		},
		{
			name: "success",
			f: func(orgSlug string) error {
				return errors.New("error")
			},
			expectedStatus:  500,
			expectedMessage: "something wrong",
		},
	}

	cl := pb.NewCommentServiceClient(conn)
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			mockUsecase.DelFunc = v.f
			resp, err := cl.Delete(context.Background(), &pb.DeleteReq{})
			assert.NoError(t, err)
			assert.Equal(t, v.expectedStatus, int(resp.GetStatus()))
		})
	}
}
