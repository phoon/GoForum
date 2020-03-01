package helper

import (
	"strconv"

	"github.com/phoon/go-forum/config"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//GenUserRelatedFields generates the fields only appear in output json data.
func GenUserRelatedFields(user *model.User) {
	user.FollowerCnt = uint(repository.UserRepo.GetUserFollowersCnt(user))
	user.FollowingCnt = uint(repository.UserRepo.GetUserFollowingCnt(user))
	user.LikedTopicsCnt = uint(repository.UserRepo.GetUserLikedTopicsCnt(user))
	user.CommentCnt = uint(repository.UserRepo.GetUserCommentsCnt(user))
	user.TopicCnt = uint(repository.UserRepo.GetUserTopicsCnt(user))

	user.FollowersURL = config.Fields.BaseURL + "/users/" + user.Name + "/followers"
	user.FollowingURL = config.Fields.BaseURL + "/users/" + user.Name + "/following"
	user.TopicsURL = config.Fields.BaseURL + "/users/" + user.Name + "/topics"
	user.LikedTopicsURL = config.Fields.BaseURL + "/users/" + user.Name + "/liked"
	user.CommentsURL = config.Fields.BaseURL + "/users/" + user.Name + "/comments"
}

//GenTopicRelatedFields do the similar works as function above.
func GenTopicRelatedFields(topic *model.Topic) {
	id := strconv.FormatUint(uint64(topic.ID), 10)
	topic.LikedCnt = uint(repository.TopicRepo.GetLikedCnt(topic))
	topic.CommentsURL = config.Fields.BaseURL + "/topics/" + id + "/comments"
}

//GenCategoryRelatedFields do the similar works like above.
func GenCategoryRelatedFields(category *model.Category) {
	id := strconv.FormatUint(uint64(category.ID), 10)
	category.TopicsURL = config.Fields.BaseURL + "/categories/" + id + "/topics"

	category.TopicCnt = uint(repository.CategoryRepo.GetTopicsCnt(category))
}
