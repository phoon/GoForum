package repository

import "github.com/phoon/go-forum/repository/model"

var CommentRepo = newCommentRepo()

type commentRepo struct{}

func newCommentRepo() *commentRepo {
	return &commentRepo{}
}

func (c *commentRepo) GetAll(limit, offset int) []*model.Comment {
	var ret []*model.Comment
	if limit > 0 && offset >= 0 {
		if err := db.Limit(limit).Offset(offset).Order("id asc").Find(&ret).Error; err != nil {
			return nil
		}
	} else {
		if err := db.Find(&ret).Error; err != nil {
			return nil
		}
	}
	return ret
}

func (c *commentRepo) FindByID(id uint) *model.Comment {
	ret := &model.Comment{}
	if err := db.First(ret, id).Error; err != nil {
		return nil
	}
	return ret
}

func (c *commentRepo) Add(comment *model.Comment) (err error) {
	err = db.Create(comment).Error
	return
}

func (c *commentRepo) Update(comment *model.Comment) (err error) {
	err = db.Model(&model.Comment{}).Where("id = ?", comment.ID).Update(comment).Error
	return
}

func (c *commentRepo) Delete(id uint) (err error) {
	err = db.Where("id = ?", id).Delete(&model.Comment{}).Error
	return
}
