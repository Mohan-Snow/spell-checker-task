package service

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"spell-checker/internal/model"
)

const spellingCheckExternalService = "https://speller.yandex.net/services/spellservice.json/checkText"

type SpellingCheckService struct {
	logger *zap.Logger
}

func NewSpellingCheckService(l *zap.Logger) *SpellingCheckService {
	return &SpellingCheckService{logger: l}
}

func (s *SpellingCheckService) CheckSpelling(text *model.Text) (string, error) {
	//values := url.Values{}
	//values.Add("text", text.Text)
	textJson, _ := json.Marshal(text)
	response, err := s.doCheckRequest(textJson)
	if err != nil {
		return "", err
	}
	s.logger.Sugar().Infof("ResponseModel: %v", response)
	return text.Text, nil
}

func (s *SpellingCheckService) doCheckRequest(textJson []byte) (*[]model.CheckedText, error) {

	request, err := http.NewRequest("POST", spellingCheckExternalService, bytes.NewBuffer(textJson))
	if err != nil {
		// TODO: handle error
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	s.logger.Sugar().Infof("Request: %v", request.Body)
	client := &http.Client{}
	response, err := client.Do(request)

	//response, err := http.PostForm(spellingCheckExternalService, values)
	defer response.Body.Close()
	if err != nil {
		// TODO: handle error
		return nil, err
	}
	//io.Copy(os.Stdout, response.Body) // output response body to stdout
	s.logger.Sugar().Infof("Response: %v", response.Body)
	var responseModel = make([]model.CheckedText, 1)
	body := json.NewDecoder(response.Body).Decode(&responseModel)
	s.logger.Sugar().Infof("Body: %s", body) // check here!!!
	return &responseModel, nil
}
