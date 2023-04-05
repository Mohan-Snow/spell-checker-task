package service

import "spell-checker/cmd/model"

type SpellingCheckService struct {
}

func NewSpellingCheckService() *SpellingCheckService {
	return &SpellingCheckService{}
}

func (s *SpellingCheckService) CheckSpelling(text *model.Text) (string, error) {
	checkedText := text.Text
	return checkedText, nil
}
