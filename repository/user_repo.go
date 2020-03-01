package repository

import "github.com/phoon/go-forum/repository/model"

var UserRepo = newUserRepo()

type userRepo struct{}

func newUserRepo() *userRepo {
	return &userRepo{}
}

func (u *userRepo) FindByID(id uint) *model.User {
	ret := &model.User{}
	if err := db.First(ret, id).Error; err != nil {
		return nil
	}
	return ret
}

func (u *userRepo) FindByUserName(name string) *model.User {
	ret := &model.User{}
	if err := db.Where("name = ?", name).First(ret).Error; err != nil {
		return nil
	}
	return ret
}

func (u *userRepo) FindByEmail(email string) *model.User {
	ret := &model.User{}
	if err := db.Where("email = ?", email).First(ret).Error; err != nil {
		return nil
	}
	return ret
}

func (u *userRepo) GetAll(limit, offset int) []*model.User {
	var ret []*model.User

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

func (u *userRepo) GetRelatedTopics(user *model.User, limit, offset int) []*model.Topic {
	var ret []*model.Topic
	if limit > 0 && offset >= 0 {
		if err := db.Model(user).Limit(limit).Offset(offset).Order("id asc").Related(&ret).Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(user).Related(&ret).Error; err != nil {
			return nil
		}
	}
	return ret
}

func (u *userRepo) GetLikedTopics(user *model.User, limit, offset int) []*model.Topic {
	var ret []*model.Topic
	if limit > 0 && offset >= 0 {
		if err := db.Model(user).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "LikedTopics").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(user).Related(&ret, "LikedTopics").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (u *userRepo) GetRelatedComments(user *model.User, limit, offset int) []*model.Comment {
	var ret []*model.Comment
	if limit > 0 && offset >= 0 {
		if err := db.Model(user).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "Comments").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(user).Related(&ret, "Comments").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (u *userRepo) GetUserFollowers(user *model.User, limit, offset int) []*model.User {
	var ret []*model.User
	if limit > 0 && offset >= 0 {
		if err := db.Model(user).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "Followers").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(user).Related(&ret, "Followers").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (u *userRepo) GetUserFollowing(user *model.User, limit, offset int) []*model.User {
	var ret []*model.User
	if limit > 0 && offset >= 0 {
		if err := db.Model(user).Limit(limit).Offset(offset).Order("id asc").Related(&ret, "Following").Error; err != nil {
			return nil
		}
	} else {
		if err := db.Model(user).Related(&ret, "Following").Error; err != nil {
			return nil
		}
	}
	return ret
}

func (u *userRepo) FollowOthers(user, user2follow *model.User) (err error, follow bool) {
	var followers []*model.User
	db.Model(user2follow).Association("followers").Find(&followers)
	for _, follower := range followers {
		if follower.ID == user.ID {
			err = db.Model(user2follow).Association("followers").Delete(user).Error
			if err != nil {
				return
			}
			err = db.Model(user).Association("Following").Delete(user2follow).Error
			return
		}
	}
	err = db.Model(user2follow).Association("Followers").Append(user).Error
	if err != nil {
		return
	}
	err = db.Model(user).Association("Following").Append(user2follow).Error
	if err != nil {
		return
	}

	return nil, true
}

func (u *userRepo) GetUserFollowingCnt(user *model.User) int {
	return db.Model(user).Association("Following").Count()
}

func (u *userRepo) GetUserFollowersCnt(user *model.User) int {
	return db.Model(user).Association("Followers").Count()
}

func (u *userRepo) GetUserTopicsCnt(user *model.User) int {
	return db.Model(user).Association("Topics").Count()
}

func (u *userRepo) GetUserCommentsCnt(user *model.User) int {
	return db.Model(user).Association("Comments").Count()
}

func (u *userRepo) GetUserLikedTopicsCnt(user *model.User) int {
	return db.Model(user).Association("LikedTopics").Count()
}

func (u *userRepo) Add(user *model.User) (err error) {
	err = db.Create(user).Error
	return
}

func (u *userRepo) Update(user *model.User) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", user.ID).Omit("name").Updates(user).Error
	return
}

func (u *userRepo) UpdateColumn(id uint, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Update(name, value).Error
	return
}

func (u *userRepo) Delete(id uint) (err error) {
	//permanently delete
	err = db.Unscoped().Where("id = ?", id).Delete(&model.User{}).Error
	return
}
