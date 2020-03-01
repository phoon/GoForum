package repository

import "github.com/phoon/go-forum/repository/model"

var TopicRepo = newTopicRepo()

type topicRepo struct{}

func newTopicRepo() *topicRepo {
	return &topicRepo{}
}

func (t *topicRepo) FindByID(id uint) *model.Topic {
	ret := &model.Topic{}
	if err := db.First(ret, id).Error; err != nil {
		return nil
	}
	return ret
}

func (t *topicRepo) GetAll(limit, offset int) []*model.Topic {
	var ret []*model.Topic

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

func (t *topicRepo) GetRelatedComments(topic *model.Topic, limit, offset int) []*model.Comment {
	var ret []*model.Comment

	if limit > 0 && offset >= 0 {
		if err := db.Model(topic).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "Comments").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(topic).Related(&ret, "Comments").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (t *topicRepo) GotLiked(topic *model.Topic, likeduser *model.User) (err error, like bool) {
	var users []*model.User
	db.Model(topic).Association("LikedBy").Find(&users)
	for _, user := range users {
		if user.ID == likeduser.ID {
			err = db.Model(topic).Association("LikedBy").Delete(likeduser).Error
			if err != nil {
				return err, false
			}
			err = db.Model(likeduser).Association("LikedTopics").Delete(topic).Error
			return
		}
	}

	err = db.Model(topic).Association("LikedBy").Append(likeduser).Error
	if err != nil {
		return
	}
	err = db.Model(likeduser).Association("LikedTopics").Append(topic).Error
	if err != nil {
		return
	}
	return nil, true
}

func (t *topicRepo) GetLikedCnt(topic *model.Topic) int {
	return db.Model(topic).Association("LikedBy").Count()
}

func (t *topicRepo) Add(topic *model.Topic) (err error) {
	err = db.Create(topic).Error
	return
}

func (t *topicRepo) Update(topic *model.Topic) (err error) {
	err = db.Model(&model.Topic{}).Where("id = ?", topic.ID).Update(topic).Error
	return
}

func (t *topicRepo) Delete(id uint) (err error) {
	err = db.Where("id = ?", id).Delete(&model.Topic{}).Error
	return
}
