package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//GetALlCategories gets all the categories, paging is support.
func GetAllCategories(ctx *gin.Context) {
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	categories := repository.CategoryRepo.GetAll(limit, offset)
	for _, category := range categories {
		helper.GenCategoryRelatedFields(category)
	}
	helper.CtxJSON(ctx, "", categories)
}

//GetCategoryByID get the category by specified id.
func GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	category := repository.CategoryRepo.FindByID(uint(id))
	helper.GenCategoryRelatedFields(category)
	helper.CtxJSON(ctx, "", category)
}

//GetCategoryTopics get all the topics that belongs to the category, paging is support.
func GetCategoryTopics(ctx *gin.Context) {
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

	category := repository.CategoryRepo.FindByID(uint(id))
	if category == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "category does not exists")
		return
	}

	topics := repository.CategoryRepo.GetRelatedTopics(category, limit, offset)
	for _, topic := range topics {
		helper.GenTopicRelatedFields(topic)
	}
	helper.CtxJSON(ctx, "", topics)
}

//CreateCategory create a new category, only super user can do this.
//Admin middleware has ben called in advance, so rest assured.
func CreateCategory(ctx *gin.Context) {
	var createCategoryFrom model.Category
	if err := ctx.ShouldBind(&createCategoryFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	if err := repository.CategoryRepo.Add(&createCategoryFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to create the category: "+err.Error())
		return
	}

	helper.GenCategoryRelatedFields(&createCategoryFrom)
	helper.CtxJSON(ctx, "create the categoty successfully", createCategoryFrom)
}

//UpdateCategory updates a category, only with super user's permission.
//Admin middleware has ben called in advance, so rest assured.
func UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	category := repository.CategoryRepo.FindByID(uint(id))
	if category == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "category does not exists")
		return
	}

	var updateCategoryFrom model.Category
	if err := ctx.ShouldBind(&updateCategoryFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	updateCategoryFrom.ID = category.ID
	if err := repository.CategoryRepo.Update(&updateCategoryFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update the category: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "update the category successfully", updateCategoryFrom)
}

//DeleteCategory deletes the given category.
//Admin middleware has ben called in advance, so rest assured.
func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	category := repository.CategoryRepo.FindByID(uint(id))
	if category == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "category does not exists")
		return
	}

	if err = repository.CategoryRepo.Delete(uint(id)); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to delete the category: "+err.Error())
	}

	helper.CtxJSON(ctx, "delete the category successfully", nil)
}
