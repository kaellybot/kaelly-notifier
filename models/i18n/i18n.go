package i18n

import (
	"embed"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/de"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/pt"
	amqp "github.com/kaellybot/kaelly-amqp"
	"golang.org/x/text/language"
)

type Language struct {
	Locale          discordgo.Locale
	Tag             language.Tag
	DateTranslator  locales.Translator
	AMQPLocale      amqp.Language
	TranslationFile string
}

const (
	frenchFile     = "fr.json"
	englishFile    = "en.json"
	spanishFile    = "es.json"
	germanFile     = "de.json"
	portugueseFile = "pt.json"

	DefaultAMQPLocale = amqp.Language_EN
	DefaultLocale     = discordgo.EnglishGB
	InternalLocale    = discordgo.French
)

//go:embed *.json
var Folder embed.FS

func GetLanguages() []Language {
	return []Language{
		{
			Locale:          discordgo.French,
			Tag:             language.French,
			DateTranslator:  fr.New(),
			TranslationFile: frenchFile,
			AMQPLocale:      amqp.Language_FR,
		},
		{
			Locale:          discordgo.EnglishGB,
			Tag:             language.English,
			DateTranslator:  en.New(),
			TranslationFile: englishFile,
			AMQPLocale:      amqp.Language_EN,
		},
		{
			Locale:          discordgo.EnglishUS,
			Tag:             language.English,
			DateTranslator:  en_US.New(),
			TranslationFile: englishFile,
			AMQPLocale:      amqp.Language_EN,
		},
		{
			Locale:          discordgo.SpanishES,
			Tag:             language.Spanish,
			DateTranslator:  es.New(),
			TranslationFile: spanishFile,
			AMQPLocale:      amqp.Language_ES,
		},
		{
			Locale:          discordgo.German,
			Tag:             language.German,
			DateTranslator:  de.New(),
			TranslationFile: germanFile,
			AMQPLocale:      amqp.Language_DE,
		},
		{
			Locale:          discordgo.PortugueseBR,
			Tag:             language.Portuguese,
			DateTranslator:  pt.New(),
			TranslationFile: portugueseFile,
			AMQPLocale:      amqp.Language_PT,
		},
	}
}

func GetLanguage(locale amqp.Language) Language {
	for _, language := range GetLanguages() {
		if language.AMQPLocale == locale {
			return language
		}
	}

	return GetLanguage(DefaultAMQPLocale)
}

func MapTag(locale discordgo.Locale) language.Tag {
	if locale == DefaultLocale {
		return language.English
	}

	for _, language := range GetLanguages() {
		if language.Locale == locale {
			return language.Tag
		}
	}

	return MapTag(DefaultLocale)
}
