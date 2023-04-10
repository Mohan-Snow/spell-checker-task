package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"spell-checker/internal/model"
)

type TextHandler struct {
	logger  *zap.Logger
	service SpellingCheckService
}

type SpellingCheckService interface {
	CheckSpelling(text *model.Text) (string, error)
}

func NewTextHandler(logger *zap.Logger, s SpellingCheckService) *TextHandler {
	return &TextHandler{logger: logger, service: s}
}

func (t *TextHandler) CheckForSpelling(context *gin.Context) {
	var input model.Text
	// ShouldBindJSON(&input) - pass pointer to structure
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	checkedText, err := t.service.CheckSpelling(&input)
	if err != nil {
		// TODO: handle error
	}
	context.JSON(http.StatusOK, gin.H{"data": checkedText})
}
