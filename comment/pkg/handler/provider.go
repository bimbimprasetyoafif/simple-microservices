package handler

import "github.com/bimbimprasetyoafif/comment/pkg/model"

type UsecaseProvider interface {
	CreateComment(orgSlug, commentValue string) (bool, error)
	GetAllComment(orgSlug string) ([]model.Comment, error)
	DeleteComment(orgSlug string) error
}
