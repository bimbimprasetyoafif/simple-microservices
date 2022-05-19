package grpc

import (
	"fmt"
	"github.com/bimbimprasetyoafif/comment/config"
	"github.com/bimbimprasetyoafif/comment/pkg/handler"
	"github.com/bimbimprasetyoafif/comment/pkg/repository"
	"github.com/bimbimprasetyoafif/comment/pkg/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	S             *grpc.Server
	DB            *gorm.DB
	Cfg           config.Config
	Ready         chan bool
	AddressListen net.Listener
	OrgClient     usecase.OrganizationClientProvider
}

func (serv *Server) RegisterServer() {
	repo := repository.NewCommentRepo(serv.DB)
	uc := usecase.NewCommentUsecase(repo, serv.OrgClient)

	handler.NewGrpcHandler(serv.S, uc)
}

func (serv *Server) Start() {
	serv.RegisterServer()

	errChan := make(chan error)
	if serv.AddressListen == nil {
		listener, err := net.Listen("tcp", serv.Cfg.ServerPort)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		serv.AddressListen = listener
	}

	go func() {
		if err := serv.S.Serve(serv.AddressListen); err != nil {
			errChan <- err
		}
	}()
	fmt.Println("server running at", serv.Cfg.ServerPort)

	if serv.Ready != nil {
		serv.Ready <- true
	}

	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM)

	<-interruptSignal
	fmt.Println(" <-- ups, server was interrupt")
	serv.S.GracefulStop()
}
