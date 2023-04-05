package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spell-checker/cmd/model"
)

type TextHandler struct {
	service SpellingCheckService
}

type TextCheckRequest struct {
	RequestString string
}

type SpellingCheckService interface {
	CheckSpelling(text *model.Text) (string, error)
}

func NewTextHandler(s SpellingCheckService) *TextHandler {
	return &TextHandler{service: s}
}

func (t *TextHandler) CheckForSpelling(context *gin.Context) {
	var input TextCheckRequest
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newText := model.Text{Text: input.RequestString}
	checkedText, err := t.service.CheckSpelling(&newText)
	if err != nil {
		// TODO: handle error
	}
	context.JSON(http.StatusOK, gin.H{"data": checkedText})
}
