package handler

import "github.com/bimbimprasetyoafif/comment/pkg/model"

type UsecaseMock struct {
	CreateFunc func(orgSlug, commentValue string) (bool, error)
	GetFunc    func(orgSlug string) ([]model.Comment, error)
	DelFunc    func(orgSlug string) error
}

func (u *UsecaseMock) CreateComment(orgSlug, commentValue string) (bool, error) {
	return u.CreateFunc(orgSlug, commentValue)
}

func (u *UsecaseMock) GetAllComment(orgSlug string) ([]model.Comment, error) {
	return u.GetAllComment(orgSlug)
}

func (u *UsecaseMock) DeleteComment(orgSlug string) error {
	return u.DelFunc(orgSlug)
}
