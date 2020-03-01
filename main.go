package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/config"
	"github.com/phoon/go-forum/handler/router"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

func main() {
	config.Load()
	repository.Start()

	gin.SetMode(config.Fields.Mode)
	gin.DisableConsoleColor()

	//set the gin framework log file
	logf, err := os.OpenFile(config.Fields.GinLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Fail to open log file for gin framework.")
		os.Exit(1)
	}
	defer logf.Close()
	gin.DefaultWriter = io.MultiWriter(logf)

	//use cookie to store session
	store := sessions.NewCookieStore([]byte(config.Fields.GinSessionKey))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   1296000, //15 dys
		HttpOnly: true,
	})
	gob.Register(&model.LoginInfo{}) //register for securecookie

	app := gin.Default()
	app.Use(sessions.Sessions("SESSION", store))
	router.ApplyRoutes(app)
	pprof.Register(app)

	app.Run(config.Fields.Address)
}
