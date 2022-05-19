package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/comment/config"
	"github.com/bimbimprasetyoafif/comment/database"
	cl "github.com/bimbimprasetyoafif/comment/pkg/client/grpc"
	grpc_server "github.com/bimbimprasetyoafif/comment/server/grpc"
	"google.golang.org/grpc"
)

func init() {
	fmt.Println("======COMMENT=====")
}

func main() {
	cfg := config.InitConfig()

	oClient, err := cl.InitOrgClient(cfg.OrgUrl, cfg.OrgPort)
	if err != nil {
		return
	}

	db, err := database.InitDatabase(cfg.DatabaseUname, cfg.DatabasePass, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server := grpc.NewServer()
	ready := make(chan bool)
	s := grpc_server.Server{
		S:         server,
		DB:        db,
		Cfg:       cfg,
		Ready:     ready,
		OrgClient: oClient,
	}

	s.Start()
}
