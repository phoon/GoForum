package repository

import "github.com/phoon/go-forum/repository/model"

var CategoryRepo = newCateGoryRepo()

type categoryRepo struct{}

func newCateGoryRepo() *categoryRepo {
	return &categoryRepo{}
}

func (c *categoryRepo) GetAll(limit, offset int) []*model.Category {
	var ret []*model.Category
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

func (c *categoryRepo) FindByID(id uint) *model.Category {
	ret := &model.Category{}
	if err := db.First(ret, id).Error; err != nil {
		return nil
	}
	return ret
}

func (c *categoryRepo) GetRelatedTopics(category *model.Category, limit, offset int) []*model.Topic {
	var ret []*model.Topic
	if limit > 0 && offset >= 0 {
		if err := db.Model(category).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "Topics").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(category).Related(&ret, "Topics").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (c *categoryRepo) GetTopicsCnt(category *model.Category) int {
	return db.Model(category).Association("Topics").Count()
}

func (c *categoryRepo) Add(category *model.Category) (err error) {
	err = db.Create(category).Error
	return
}

func (c *categoryRepo) Update(category *model.Category) (err error) {
	err = db.Model(&model.Category{}).Where("id = ?", category.ID).Update(category).Error
	return
}

func (c *categoryRepo) Delete(id uint) (err error) {
	err = db.Where("id = ?", id).Delete(&model.Category{}).Error
	return
}
