package handler

import (
	"github.com/bimbimprasetyoafif/api-gateway/pkg/handler/comment"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/handler/organization"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"github.com/labstack/echo/v4"
)

func RegisterOrganization(e *echo.Echo, provider provider.OrgClientProvider) {
	c := organization.Controller{
		Cl: provider,
	}

	e.POST("/orgs", c.RegisterOrganization)
}

func RegisterComment(e *echo.Echo, provider provider.CommentClientProvider) {
	c := comment.Controller{
		Cl: provider,
	}

	e.POST("/orgs/:slug/comments", c.CreateComment)
	e.GET("/orgs/:slug/comments", c.GetAllComment)
	e.DELETE("/orgs/:slug/comments", c.DeleteComment)
}
