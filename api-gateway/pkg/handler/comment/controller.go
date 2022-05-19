package comment

import (
	"fmt"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Cl provider.CommentClientProvider
}

func (o *Controller) CreateComment(ctx echo.Context) error {
	body := struct {
		Comment string `json:"comment"`
	}{}

	ctx.Bind(&body)
	res, err := o.Cl.CreateComment(ctx.Param("slug"), body.Comment)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "internal error",
		})
	}

	return ctx.JSON(int(res.Status), map[string]interface{}{
		"status":  res.GetStatus(),
		"message": res.GetMessage(),
	})
}

func (o *Controller) GetAllComment(ctx echo.Context) error {
	res, err := o.Cl.GetAllComment(ctx.Param("slug"))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "internal error",
		})
	}

	return ctx.JSON(int(res.Status), map[string]interface{}{
		"status":  res.GetStatus(),
		"message": res.GetMessage(),
		"data":    res.GetData(),
	})
}

func (o *Controller) DeleteComment(ctx echo.Context) error {
	res, err := o.Cl.DeleteComment(ctx.Param("slug"))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "internal error",
		})
	}

	return ctx.JSON(int(res.Status), map[string]interface{}{
		"status":  res.GetStatus(),
		"message": res.GetMessage(),
	})
}
