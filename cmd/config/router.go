package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spell-checker/cmd/handler"
	"spell-checker/cmd/service"
)

type Router struct {
	textHandler TextHandler
}

type TextHandler interface {
	CheckForSpelling(context *gin.Context)
}

func (r *Router) SetupRouter() *gin.Engine {
	ginRouter := gin.Default()
	s := service.NewSpellingCheckService()
	r.textHandler = handler.NewTextHandler(s)

	// ping test
	ginRouter.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// setup checkSpelling endpoint
	ginRouter.POST("/checkSpelling", r.textHandler.CheckForSpelling) // how context is passed here?
	return ginRouter
}
