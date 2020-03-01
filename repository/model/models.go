package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	UserRoleRegular = 0
	UserRoleAdmin   = 1

	UserStatusNormal   = 0
	UserStatusInactive = 1
	UserStatusBanned   = 2

	TopicStatusNormal = 0
	TopicStatusBanned = 1

	CommentStatusNormal = 0
	CommentStatusBanned = 1
)

var (
	//Models is used for auto migration.
	Models = []interface{}{
		&User{},
		&Topic{},
		&Category{},
		&Comment{},
	}
)

//BasicField Model
type BasicField struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

//User Model
type User struct {
	BasicField
	//Name is a unique username, up to 15 bytes
	Name string `json:"username" form:"username" binding:"required" gorm:"size:16;unique_index;not null"`
	//Nickname can be up to 15 bytes
	Nickname string `json:"nickname" form:"nickname" gorm:"size:21"`
	//Email is a unique email address, up to 128 bytes
	Email string `json:"email" form:"email" binding:"required" gorm:"size:129;unique_index;not null"`
	//Password with bcrypt processed, 60 bytes long
	Password string `json:"-" form:"password" binding:"required" gorm:"type:char(60)"`
	//Bio is the user's own description, up to 120 bytes
	Bio string `json:"bio" form:"bio" gorm:"size:121;default:'这个用户很懒，什么都没留下'"`
	//Role indentifies the user's indentity
	Role uint8 `json:"role"`
	//AvatarURL is user's avatar url
	AvatarURL string `json:"avatar_url" form:"avatar_url"`

	//Data statistics
	TopicCnt       uint `json:"topics_count" gorm:"-"`
	CommentCnt     uint `json:"comments_count" gorm:"-"`
	LikedTopicsCnt uint `json:"liked_topics_count" gorm:"-"`
	FollowerCnt    uint `json:"followers_count" gorm:"-"`
	FollowingCnt   uint `json:"following_count" gorm:"-"`

	//Associations
	Followers   []*User   `json:"-" gorm:"many2many:user_followers;association_jointable_foreignkey:follower_id"`
	Following   []*User   `json:"-" gorm:"many2many:user_following;association_jointable_foreignkey:following_id"`
	LikedTopics []*Topic   `json:"-" gorm:"many2many:user_liked_topics"`
	Topics      []*Topic   `json:"-" gorm:"foreignkey:UserID"`
	Comments    []*Comment `json:"-" gorm:"foreignkey:UserID"`

	//Related API urls
	FollowersURL   string `json:"followers_url" gorm:"-"`
	FollowingURL   string `json:"following_url" gorm:"-"`
	TopicsURL      string `json:"topics_url" gorm:"-"`
	LikedTopicsURL string `json:"liked_topics_url" gorm:"-"`
	CommentsURL    string `json:"comments_url" gorm:"-"`

	//Status of user: 0 - normal, 1 - email inactive, 2 - banned
	Status uint8 `json:"status" gorm:"index:idx_status;not null"`
}

//Topic Model
type Topic struct {
	BasicField
	//Title of the topic, up to 60 bytes
	Title string `json:"title" form:"title" binding:"required" gorm:"size:61;not null"`
	//Summary of the topic
	Summary string `json:"summary" form:"summary" binding:"required"`
	//UserID foreignkey
	UserID uint `json:"user_id"`
	//CategoryID foreignkey
	CategoryID uint `json:"category_id" form:"category_id" binding:"required"`
	//RawContent is the original markdown content
	RawConent string `json:"markdown" form:"markdown" binding:"required" gorm:"type:longtext"`
	//HTMLContent is the generated HTML content from markdown
	HTMLContent string `json:"html" form:"html" binding:"required" gorm:"type:longtext"`
	//LikedCnt means how many people hit the like button
	LikedCnt uint `json:"liked_count" gorm:"-"`

	//Association
	Comments []*Comment `json:"-" gorm:"foreignkey:TopicID"`
	LikedBy  []*User    `json:"-" gorm:"many2many:user_liked_topics"`

	CommentsURL string `json:"comments_url" gorm:"-"`

	//Status of topic: 0 - normal, 1 - banned
	Status uint8 `json:"status" gorm:"index:idx_status;not null"`
}

//Category Model
type Category struct {
	BasicField
	//Name of the category, up to 200 bytes
	Name string `json:"name" form:"name" binding:"required" gorm:"size:21;unique;not null"`
	//Description of the category, up to 100 bytes
	Description string `json:"description" form:"description" gorm:"size:101;default:'暂无描述'"`
	//TopicCnt means how many topics under this category
	TopicCnt uint `json:"topic_count" gorm:"-"`

	TopicsURL string `json:"topics_url" gorm:"-"`

	//Topics that related to this category, foreign key
	Topics []*	Topic `json:"-" gorm:"foreignkey:CategoryID"`

	//Status of category:TODO
	Status uint8 `gorm:"index:idx_status;not null"`
}

//Comment Model
type Comment struct {
	BasicField
	//Content of the comment
	Content string `json:"content" form:"content" binding:"required" gorm:"type:text;not null"`

	//UserID means which user is this comment belongs to
	UserID uint `json:"user_id"`
	//TopicID means which topic is this comment belongs to
	TopicID uint `json:"topic_id" form:"topic_id" binding:"required"`
	//QuoteID means which comment is this comment quotes to
	QuoteID uint `json:"quote_id" form:"quote_id"`

	//Status of comment: 0 - normal, 1 - banned
	Status uint8 `gorm:"index:idx_status;not null"` //评论状态: 0 正常; 1 被封禁
}

//UserLoginInput Modle
type UserLoginInput struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//CustomClaim Model
type CustomClaim struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	jwt.StandardClaims
}

//LoginInfo
type LoginInfo struct {
	UserID        uint
	UserName      string
	Role          uint8
	Authenticated bool
}
