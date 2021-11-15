package model

import "golang.org/x/text/language"

type Language string

func (l Language) String() string {
	return string(l)
}

func (l Language) GetLocale() language.Tag {
	switch l {
	case English:
		return language.English
	case Russian:
		return language.Russian
	case Ukrainian:
		return language.Ukrainian
	case Polish:
		return language.Polish
	}
	return language.English
}

const (
	English         Language = "en"
	Russian         Language = "ru"
	Ukrainian       Language = "ua"
	Polish          Language = "pl"
	UnknownLanguage Language = "unknown"
)

func EncodeLanguage(input string) Language {
	switch input {
	case English.String():
		return English
	case Russian.String():
		return Russian
	case Ukrainian.String():
		return Ukrainian
	case Polish.String():
		return Polish

	default:
		return UnknownLanguage
	}
}
