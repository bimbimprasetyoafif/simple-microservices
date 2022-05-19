package usecase

import (
	"fmt"
	"github.com/bimbimprasetyoafif/comment/pkg/model"
)

type commentUsecase struct {
	repo CommentProvider
	cl   OrganizationClientProvider
}

func (c *commentUsecase) CreateComment(orgSlug, commentValue string) (bool, error) {
	resClient, err := c.cl.CheckOrg(orgSlug)
	if err != nil {
		fmt.Println("error call grpc server: ", err)
		return false, err
	}

	if !resClient.IsExist {
		return false, nil
	}

	if err := c.repo.InsertOne(int(resClient.OrgId), commentValue); err != nil {
		return false, err
	}

	return true, nil
}

func (c *commentUsecase) GetAllComment(orgSlug string) ([]model.Comment, error) {
	resClient, err := c.cl.CheckOrg(orgSlug)
	if err != nil {
		fmt.Println("error call grpc server: ", err)
		return nil, err
	}

	if !resClient.IsExist {
		return nil, nil
	}

	return c.repo.GetAllByOrg(int(resClient.OrgId))
}

func (c *commentUsecase) DeleteComment(orgSlug string) error {
	resClient, err := c.cl.CheckOrg(orgSlug)
	if err != nil {
		fmt.Println("error call grpc server: ", err)
		return err
	}

	if !resClient.IsExist {
		return nil
	}

	return c.repo.DeleteByOrg(int(resClient.OrgId))
}

func NewCommentUsecase(p CommentProvider, cl OrganizationClientProvider) *commentUsecase {
	return &commentUsecase{
		repo: p,
		cl:   cl,
	}
}
