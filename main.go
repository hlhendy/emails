package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/emails/internal/messages"
	"github.com/julienschmidt/httprouter"
)

func main() {
	settings, err := getSettings("settings.json")
	if err != nil {
		log.Fatalf("error opening settings: %+v", err)
	}

	if settings == nil {
		log.Fatal("settings are nil")
	}

	var defaultSettings provider
	for _, p := range settings.Providers {
		if p.Name == settings.DefaultProvider {
			defaultSettings = p
			break
		}
	}

	message, err := messages.New(settings.DefaultProvider, defaultSettings.APIKey, defaultSettings.BaseURL)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	h := NewHandler(message)

	router := httprouter.New()
	router.POST("/email", h.Send)

	log.Printf("listening on port %s", settings.Port)
	log.Fatal(http.ListenAndServe(settings.Port, router))
}

type settings struct {
	DefaultProvider string     `json:"default_provider"`
	ValidProviders  []string   `json:"valid_providers"`
	Port            string     `json:"port"`
	LogLevel        string     `json:"log_level"`
	Providers       []provider `json:"providers"`
}

type provider struct {
	Name    string `json:"name"`
	BaseURL string `json:"base_url"`
	APIKey  string `json:"api_key"`
}

func getSettings(settingsFile string) (*settings, error) {
	sf, err := os.Open(settingsFile)
	if err != nil {
		return nil, err
	}

	defer sf.Close()

	body, err := ioutil.ReadAll(sf)
	if err != nil {
		return nil, err
	}

	var s settings
	if err := json.Unmarshal(body, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
