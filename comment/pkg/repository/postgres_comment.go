package repository

import (
	"github.com/bimbimprasetyoafif/comment/pkg/model"
	"gorm.io/gorm"
)

type postgresComment struct {
	db *gorm.DB
}

func (p *postgresComment) InsertOne(orgID int, commentValue string) error {
	var comment model.Comment
	comment.OrgID = orgID
	comment.Value = commentValue

	return p.db.Create(&comment).Error
}

func (p *postgresComment) GetAllByOrg(orgID int) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)

	err := p.db.Where("org_id = ?", orgID).Find(&comments).Error
	return comments, err
}

func (p *postgresComment) DeleteByOrg(orgID int) error {
	return p.db.Where("org_id = ?", orgID).Delete(&model.Comment{}).Error
}

func NewCommentRepo(db *gorm.DB) *postgresComment {
	return &postgresComment{
		db: db,
	}
}
