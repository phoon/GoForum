package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/middleware"
	"github.com/phoon/go-forum/handler"
)

//ApplyRoutes set the routers
func ApplyRoutes(rtr *gin.Engine) {
	//authtication related
	authAPI := rtr.Group("/auth")
	{
		authAPI.POST("/signup", handler.SignUp)
		authAPI.POST("/signin", handler.SignIn)
		authAPI.POST("/signout", middleware.LoginNeed(), handler.SignOut)
	}

	//user related
	userAPI := rtr.Group("/users")
	{
		userAPI.GET("", handler.GetAllUsers)
		userAPI.GET("/:name", handler.GetUserByName)
		userAPI.	PUT("/:name", middleware.LoginNeed(), handler.UpdateUser)
		userAPI.DELETE("/:name", middleware.LoginNeed(), handler.DeleteUser)
		userAPI.POST("/:name/follow", middleware.LoginNeed(), handler.FollowOthers)
		userAPI.GET("/:name/followers", handler.GetUserFollowers)
		userAPI.GET("/:name/following", handler.GetUserFollowing)
		userAPI.GET("/:name/topics", handler.GetUserTopics)
		userAPI.GET("/:name/comments", handler.GetUserComments)
		userAPI.GET("/:name/liked", handler.GetUserLikedTopics)
		userAPI.PUT("/:name/password", middleware.LoginNeed(), handler.UpdateUserPass)
		userAPI.PUT("/:name/email", middleware.LoginNeed(), handler.UpdateUserEmail)
		userAPI.PUT("/:name/avatar", middleware.LoginNeed(), handler.UpdateUserAvatar)
		userAPI.PUT("/:name/status", middleware.Admin(), handler.UpdateUserStatus)
	}

	//topics related
	topicAPI := rtr.Group("/topics")
	{
		topicAPI.GET("", handler.GetAllTopics)
		topicAPI.POST("", middleware.LoginNeed(), handler.CreateTopic)
		topicAPI.GET("/:id", handler.GetTopicByID)
		topicAPI.PUT("/:id", middleware.LoginNeed(), handler.UpdateTopic)
		topicAPI.DELETE("/:id", middleware.LoginNeed(), handler.DeleteTopic)
		topicAPI.GET("/:id/comments", handler.GetTopicComments)
		topicAPI.POST("/:id/like", middleware.LoginNeed(), handler.LikeTopic)
	}

	//categories related
	categoryAPI := rtr.Group("/categories")
	{
		categoryAPI.GET("", handler.GetAllCategories)
		categoryAPI.POST("", middleware.Admin(), handler.CreateCategory)
		categoryAPI.GET("/:id", handler.GetCategoryByID)
		categoryAPI.PUT("/:id", middleware.Admin(), handler.UpdateCategory)
		categoryAPI.DELETE("/:id", middleware.Admin(), handler.DeleteCategory)
		categoryAPI.GET("/:id/topics", handler.GetCategoryTopics)
	}

	//comments related
	commentAPI := rtr.Group("/comments")
	{
		commentAPI.GET("", handler.GetAllComments)
		commentAPI.POST("", middleware.LoginNeed(), handler.CreateComment)
		commentAPI.GET("/:id", handler.GetCommentByID)
		commentAPI.PUT("/:id", middleware.LoginNeed(), handler.UpdateComment)
		commentAPI.DELETE("/:id", middleware.LoginNeed(), handler.DeleteComment)
	}
}
