package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"spell-checker/internal/model"
)

const spellingCheckExternalService = "https://speller.yandex.net/services/spellservice.json/checkTexts"

type SpellingCheckService struct {
	logger *zap.Logger
}

func NewSpellingCheckService(l *zap.Logger) *SpellingCheckService {
	return &SpellingCheckService{logger: l}
}

func (s *SpellingCheckService) CheckSpelling(text *model.Text) (*model.FixedTextResponse, error) {
	//textJson, _ := json.Marshal(text)
	response, err := s.doCheckRequest(text)
	if err != nil {
		return nil, err
	}
	s.logger.Sugar().Infof("ResponseModel: %v", *response)
	var fixedText model.FixedTextResponse
	for _, val := range *response {
		for _, checkedResponse := range val {
			fixedText.InitialText = checkedResponse.InitialWord
			fixedText.FixedText = checkedResponse.FixedSpellingWords
		}
	}
	return &fixedText, nil
}

func (s *SpellingCheckService) doCheckRequest(text *model.Text) (*[]model.CheckedText, error) {
	//request, err := http.NewRequest("GET", spellingCheckExternalService, bytes.NewBuffer(textJson))
	//if err != nil {
	//	TODO: handle error
	//return nil, err
	//}
	//request.Header.Add("Content-Type", "application/json")
	//s.logger.Sugar().Infof("Request: %v", request.Body)
	//client := &http.Client{}
	//response, err := client.Do(request)

	formData := url.Values{}
	formData.Add("text", text.Text)
	response, err := http.PostForm(spellingCheckExternalService, formData)
	defer response.Body.Close()
	if err != nil {
		// TODO: handle error
		return nil, err
	}
	s.logger.Sugar().Infof("Response: %v", &response.Body)
	var responseModel []model.CheckedText
	err = json.NewDecoder(response.Body).Decode(&responseModel) // change to unmarshall
	return &responseModel, nil
}
