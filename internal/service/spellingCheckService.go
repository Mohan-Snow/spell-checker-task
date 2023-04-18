package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"

	"spell-checker/internal/model"
)

const spellingCheckExternalService = "https://speller.yandex.net/services/spellservice.json/checkTexts"

type SpellingCheckService struct {
	logger *zap.Logger
}

func NewSpellingCheckService(l *zap.Logger) *SpellingCheckService {
	return &SpellingCheckService{logger: l}
}

func (s *SpellingCheckService) CheckSpelling(initialText *model.Text) (*string, error) {
	splitText := strings.Split(initialText.Text, " ")
	checkedText, err := s.checkSplitText(splitText)
	if err != nil {
		return &initialText.Text, err
	}
	fixedText := strings.Join(*checkedText, " ")
	return &fixedText, nil
}

func (s *SpellingCheckService) checkSplitText(splitText []string) (*[]string, error) {
	response, err := s.doRequestToCheckWords(&splitText)
	if err != nil {
		return &splitText, err
	}
	s.logger.Sugar().Infof("ResponseModel: %v", *response)
	for index, sliceWithFix := range *response {
		if len(sliceWithFix) != 0 {
			splitText[index] = sliceWithFix[0].FixedSpellingWords[0]
		}
	}
	return &splitText, nil
}

func (s *SpellingCheckService) doRequestToCheckWords(splitText *[]string) (*[]model.CheckedText, error) {
	formData := url.Values{}
	for _, word := range *splitText {
		formData.Add("text", word)
	}
	response, err := http.PostForm(spellingCheckExternalService, formData)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	var responseModel []model.CheckedText
	err = json.NewDecoder(response.Body).Decode(&responseModel)
	return &responseModel, nil
}
