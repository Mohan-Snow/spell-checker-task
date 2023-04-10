package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"spell-checker/internal/handler"
	"spell-checker/internal/service"
)

type Bootstrap struct {
	textHandler TextHandler
}

type TextHandler interface {
	CheckForSpelling(context *gin.Context)
}

func (b *Bootstrap) Setup(logger *zap.Logger) *gin.Engine {
	checkService := service.NewSpellingCheckService(logger)
	b.textHandler = handler.NewTextHandler(logger, checkService)

	ginRouter := gin.Default()
	// ping test
	ginRouter.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// setup checkSpelling endpoint
	ginRouter.POST("/checkSpelling", b.textHandler.CheckForSpelling) // how context is passed here?
	return ginRouter
}
