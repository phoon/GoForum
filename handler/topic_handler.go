package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//GetAllTopics gets all the topics and paging is possible.
func GetAllTopics(ctx *gin.Context) {
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	topics := repository.TopicRepo.GetAll(limit, offset)
	for _, topic := range topics {
		helper.GenTopicRelatedFields(topic)
	}

	helper.CtxJSON(ctx, "", topics)
}

//GetTopicByID find topic by the specified topic id.
func GetTopicByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	topic := repository.TopicRepo.FindByID(uint(id))
	if topic == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "topic does not exists")
		return
	}

	helper.GenTopicRelatedFields(topic)
	helper.CtxJSON(ctx, "", topic)
}

//DeleteTopic delete a topic by specified topic id.
//
//Only the topic owner or super user(which role is model.UserRoleAdmin) can do this
func DeleteTopic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//check the given topic
	topic := repository.TopicRepo.FindByID(uint(id))
	if topic == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "topic does not exists")
		return
	}

	//check the current user's permission
	if topic.UserID != loginInfo.UserID && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to delete it")
		return
	}

	//do the delete
	if err = repository.CategoryRepo.Delete(uint(id)); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.CtxJSON(ctx, "topic was removed successfully", nil)
}

//GetTopicComments gets the comments that belongs to the topic. And paging is possible.
func GetTopicComments(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	topic := repository.TopicRepo.FindByID(uint(id))
	if topic == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "topics does not exists")
		return
	}

	comments := repository.TopicRepo.GetRelatedComments(topic, limit, offset)
	helper.CtxJSON(ctx, "", comments)
}

//CreateTopic create a new topic.
//Token is always needed, so user id should be obtained from the `claims`.
//By doing this, it can avoid forging requests and simplify the process.
func CreateTopic(ctx *gin.Context) {
	var createTopicForm model.Topic
	if err := ctx.ShouldBind(&createTopicForm); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//set topic owner
	createTopicForm.UserID = loginInfo.UserID
	if err := repository.TopicRepo.Add(&createTopicForm); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to create the topic: "+err.Error())
		return
	}

	helper.GenTopicRelatedFields(&createTopicForm)
	helper.CtxJSON(ctx, "create the topic successfully", createTopicForm)
}

//UpdateTopic updates the given topic
//
//Only the topic owner or super user can do this
func UpdateTopic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//get the topic
	topic := repository.TopicRepo.FindByID(uint(id))
	if topic == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "topic does not exists, maybe you should create it first")
		return
	}

	//check the permission
	if loginInfo.UserID != topic.UserID && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	//check the input form
	var updateTopicForm model.Topic
	if err := ctx.ShouldBind(&updateTopicForm); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	//set the owner
	updateTopicForm.ID = topic.ID
	updateTopicForm.UserID = loginInfo.UserID
	if err := repository.TopicRepo.Update(&updateTopicForm); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update the topic: "+err.Error())
		return
	}

	helper.GenTopicRelatedFields(&updateTopicForm)
	helper.CtxJSON(ctx, "update the topic successfully", updateTopicForm)
}

//LikeTopic
func LikeTopic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	topic := repository.TopicRepo.FindByID(uint(id))
	if topic == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "topic does not exists")
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	user := repository.UserRepo.FindByID(loginInfo.UserID)

	err, like := repository.TopicRepo.GotLiked(topic, user)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to like the topic: "+err.Error())
		return
	}

	if !like {
		helper.CtxJSON(ctx, "unliked successfully", nil)
		return
	}

	helper.CtxJSON(ctx, "like the topic successfully", nil)
}
