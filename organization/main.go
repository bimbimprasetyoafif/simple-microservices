package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/organization/config"
	"github.com/bimbimprasetyoafif/organization/database"
	grpc_server "github.com/bimbimprasetyoafif/organization/server/grpc"
	"google.golang.org/grpc"
)

func init() {
	fmt.Println("=====ORGANIZATION====")
}

func main() {
	cfg := config.InitConfig()
	db, err := database.InitDatabase(cfg.DatabaseUname, cfg.DatabasePass, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server := grpc.NewServer()
	ready := make(chan bool)
	s := grpc_server.Server{
		S:     server,
		DB:    db,
		Cfg:   cfg,
		Ready: ready,
	}

	s.Start()
}
