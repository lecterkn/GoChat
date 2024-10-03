package language

import "fmt"

type Language int

const (
	Undefined = iota
	Japanese
	English
	Chinese
)

type attr struct {
	V         Language
	Code      string
	Name      string
	TableName string
}

const size = 4

var languages = [size]attr{
	{Undefined, "undefined", "undefined", "undefined"},
	{Japanese, "ja", "japanese", "message_japanese_contents"},
	{English, "en", "english", "message_english_contents"},
	{Chinese, "cn", "chinese", "message_chinese_contents"},
}

func (l Language) ToCode() string {
	for _, lang := range languages {
		if lang.V == l {
			return lang.Code
		}
	}
	return ""
}

func (l Language) ToName() string {
	for _, lang := range languages {
		if lang.V == l {
			return lang.Name
		}
	}
	return ""
}

func (lang Language) TableName() string {
	return languages[lang].TableName
}

func GetLanguageFromCode(code string) (Language, error) {
	for _, lang := range languages {
		if lang.Code == code {
			return lang.V, nil
		}
	}
	return Undefined, fmt.Errorf("invalid language code")
}
