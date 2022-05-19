package testing

import (
	"context"
	"github.com/bimbimprasetyoafif/comment/config"
	"github.com/bimbimprasetyoafif/comment/database"
	pb "github.com/bimbimprasetyoafif/comment/pkg/proto"
	grpc_server "github.com/bimbimprasetyoafif/comment/server/grpc"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"os"
	"syscall"
)

const bufSize = 1024 * 1024

type e2eTestSuite struct {
	suite.Suite
	clientConn *grpc.ClientConn
	client     pb.CommentServiceClient
	lis        *bufconn.Listener
}

func (e *e2eTestSuite) bufDialer(context.Context, string) (net.Conn, error) {
	return e.lis.Dial()
}

func (e *e2eTestSuite) SetupSuite() {
	e.lis = bufconn.Listen(bufSize)

	cfg := config.InitConfig()

	clientMock := &ClientOrgMock{}
	conn, err := database.InitDatabase(cfg.DatabaseUname, cfg.DatabasePass, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
	e.Require().NoError(err)

	server := grpc.NewServer()

	ready := make(chan bool)
	s := grpc_server.Server{
		S:             server,
		DB:            conn,
		Cfg:           cfg,
		Ready:         ready,
		AddressListen: e.lis,
		OrgClient:     clientMock,
	}

	go s.Start()
	<-ready
}

func (e *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (e *e2eTestSuite) SetupTest() {
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(e.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	e.Require().NoError(err)

	e.clientConn = conn
	e.client = pb.NewCommentServiceClient(conn)
}

func (e *e2eTestSuite) TearDownTest() {
	err := e.clientConn.Close()
	e.Require().NoError(err)
}

func (e *e2eTestSuite) TestCommentCreateAndGet() {
	ctx := context.Background()

	resCreate, err := e.client.Create(ctx, &pb.CreateReq{
		OrgSlug: "surabaya-py-org",
		Value:   "first comment",
	})
	e.Require().NoError(err)
	e.Equal(200, int(resCreate.GetStatus()))

	resGet, err := e.client.GetAll(ctx, &pb.GetAllReq{
		OrgSlug: "surabaya-py-org",
	})
	e.Require().NoError(err)
	e.Equal(200, int(resGet.GetStatus()))
	resData := make([]string, 0)
	resData = append(resData, "first comment")
	e.Equal(resData, resGet.GetData())
}

func (e *e2eTestSuite) TestCommentDeleteAndGet() {
	ctx := context.Background()

	res, err := e.client.Delete(ctx, &pb.DeleteReq{
		OrgSlug: "surabaya-py-org",
	})
	e.Require().NoError(err)
	e.Equal(200, int(res.GetStatus()))

	resGet, err := e.client.GetAll(ctx, &pb.GetAllReq{
		OrgSlug: "surabaya-py-org",
	})
	e.Require().NoError(err)
	e.Equal(200, int(resGet.GetStatus()))
	e.Nil(resGet.GetData())
}
