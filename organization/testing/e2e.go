package testing

import (
	"context"
	"github.com/bimbimprasetyoafif/organization/config"
	"github.com/bimbimprasetyoafif/organization/database"
	pb "github.com/bimbimprasetyoafif/organization/pkg/proto"
	grpc_server "github.com/bimbimprasetyoafif/organization/server/grpc"
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
	client     pb.OrganizationServiceClient
	lis        *bufconn.Listener
}

func (e *e2eTestSuite) bufDialer(context.Context, string) (net.Conn, error) {
	return e.lis.Dial()
}

func (e *e2eTestSuite) SetupSuite() {
	e.lis = bufconn.Listen(bufSize)

	cfg := config.InitConfig()
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
	e.client = pb.NewOrganizationServiceClient(conn)
}

func (e *e2eTestSuite) TearDownTest() {
	err := e.clientConn.Close()
	e.Require().NoError(err)
}

func (e *e2eTestSuite) TestOrganizationCreate() {
	res, err := e.client.Create(context.Background(), &pb.CreateReq{
		Name: "surabaya.py org",
	})
	e.Require().NoError(err)

	e.Equal(200, int(res.Status))
	e.Equal("success", res.Message)
	e.Equal("surabaya-py-org", res.Data.Slug)
}

func (e *e2eTestSuite) TestOrganizationCreateExist() {
	res, err := e.client.Create(context.Background(), &pb.CreateReq{
		Name: "surabaya.py org",
	})
	e.Require().NoError(err)

	e.Equal(400, int(res.Status))
	e.Equal("org already exist", res.Message)
}

func (e *e2eTestSuite) TestOrganizationNoExist() {
	res, err := e.client.CheckAvaibility(context.Background(), &pb.CheckReq{
		Slug: "surabaya-org-com",
	})
	e.Require().NoError(err)

	e.Equal(false, res.IsExist)
	e.Equal(0, int(res.OrgId))
}

func (e *e2eTestSuite) TestOrganizationExist() {
	res, err := e.client.CheckAvaibility(context.Background(), &pb.CheckReq{
		Slug: "surabaya-py-org",
	})
	e.Require().NoError(err)

	e.Equal(true, res.IsExist)
	e.Equal(1, int(res.OrgId))
}
