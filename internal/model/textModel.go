package model

type Text struct {
	Text string `json:"text" binding:"required"`
}

type CheckedText []struct {
	Code               int      `json:"code"`
	Position           int      `json:"pos"`
	Row                int      `json:"row"`
	Col                int      `json:"col"`
	Length             int      `json:"len"`
	InitialWord        string   `json:"word"`
	FixedSpellingWords []string `json:"s"`
}
