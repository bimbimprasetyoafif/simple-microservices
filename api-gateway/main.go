package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/api-gateway/config"
	clComment "github.com/bimbimprasetyoafif/api-gateway/pkg/client/comment"
	clOrganization "github.com/bimbimprasetyoafif/api-gateway/pkg/client/organization"

	"github.com/bimbimprasetyoafif/api-gateway/pkg/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	fmt.Println("==== gateway ====")
}

func main() {
	cfg := config.InitConfig()

	clOrg, err := clOrganization.InitClient(cfg.OrgServerUrl, cfg.OrgServerPort)
	if err != nil {
		return
	}

	clCOm, err := clComment.InitClient(cfg.CommentServerUrl, cfg.CommentServerPort)
	if err != nil {
		return
	}

	h := echo.New()
	h.Use(middleware.Logger())

	handler.RegisterOrganization(h, clOrg)
	handler.RegisterComment(h, clCOm)

	h.Logger.Fatal(h.Start(cfg.ServerPort))
}
