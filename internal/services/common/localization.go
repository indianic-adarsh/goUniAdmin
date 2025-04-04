package localization

import (
	"encoding/json"
	"log"
	"path/filepath"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Bundle holds the i18n configuration
var bundle *i18n.Bundle
var once sync.Once

// Init initializes the localization bundle
func Init() {
	once.Do(func() {
		bundle = i18n.NewBundle(language.English) // Default language

		// Register a custom unmarshal function for flat key-value JSON
		bundle.RegisterUnmarshalFunc("json", func(data []byte, v interface{}) error {
			// Unmarshal into a map[string]string
			var raw map[string]string
			if err := json.Unmarshal(data, &raw); err != nil {
				return err
			}

			// Convert to the expected interface{} type (map[string]interface{})
			messages, ok := v.(*map[string]interface{})
			if !ok {
				return nil // Or handle error if needed
			}

			// Populate the messages map
			for id, value := range raw {
				(*messages)[id] = value
			}
			return nil
		})

		// Load translation files
		files, err := filepath.Glob("locales/*.json")
		if err != nil {
			log.Fatalf("Failed to find translation files: %v", err)
		}
		for _, file := range files {
			if _, err := bundle.LoadMessageFile(file); err != nil {
				log.Printf("Failed to load %s: %v", file, err)
			} else {
				log.Printf("Loaded translation file: %s", file)
			}
		}
	})
}

// GetLocalizer returns a localizer for the given language
func GetLocalizer(lang string) *i18n.Localizer {
	if bundle == nil {
		Init()
	}
	return i18n.NewLocalizer(bundle, lang)
}

// Localize retrieves a localized message
func Localize(lang, messageID string, templateData map[string]interface{}) string {
	localizer := GetLocalizer(lang)
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		log.Printf("Localization error for %s: %v", messageID, err)
		return messageID // Fallback to messageID if localization fails
	}
	return msg
}
