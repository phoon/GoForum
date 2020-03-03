package helper

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/repository/model"
	"golang.org/x/crypto/bcrypt"
)

//GenHashedPass generates the password with bcrypt processed.
func GenHashedPass(pass string) string {
	enc_pass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(enc_pass)
}

//IsPassHashed checks whether the password has been processed with bcrypt.
//It is only for the condition that a encrypted password was generated using default cost.
func IsPassHashed(pass string) bool {
	cost, err := bcrypt.Cost([]byte(pass))
	if err != nil || cost != bcrypt.DefaultCost {
		return false
	}
	return true
}

//IsEmailValid checks whether the string is a valid email address.
func IsEmailValid(email string) bool {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}

//CalLimitOffset calculates the limit and offset for database query from url params `num` and `page`.
func CalLimitOffset(ctx *gin.Context) (limit, offset int, err error) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if page <= 0 {
		err = errors.New("page should start from 1")
		return
	}
	limit, err = strconv.Atoi(ctx.DefaultQuery("num", "0"))
	if limit < 0 {
		err = errors.New("num should not be negative")
		return
	}
	offset = limit * (page - 1)
	return
}

func GetLoginInfoFromSession(ctx *gin.Context) (*model.LoginInfo, sessions.Session) {
	sess := sessions.Default(ctx)
	loginInfoField := sess.Get("login_info")
	loginInfo, ok := loginInfoField.(*model.LoginInfo)
	if !ok || loginInfo == nil {
		return nil, sess
	}

	return loginInfo, sess
}

//IsUserPassValid checks whether the password is right.
func IsUserPassValid(u *model.User, passwd string) bool {
	//user.Password has been processed with bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwd))
	if err == nil {
		return true
	}
	return false
}
