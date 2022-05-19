package organization

import (
	"fmt"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Cl provider.OrgClientProvider
}

func (o *Controller) RegisterOrganization(ctx echo.Context) error {
	body := struct {
		Name string `json:"name"`
	}{}

	ctx.Bind(&body)
	res, err := o.Cl.RegisterOrganization(body.Name)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "internal error",
		})
	}

	var data interface{} = nil

	if int(res.GetStatus()) == 200 {
		data = map[string]string{
			"name": res.Data.GetName(),
			"slug": res.Data.GetSlug(),
		}
	}

	return ctx.JSON(int(res.Status), map[string]interface{}{
		"status":  res.GetStatus(),
		"message": res.GetMessage(),
		"data":    data,
	})
}
