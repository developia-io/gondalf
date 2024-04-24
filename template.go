package main

import (
	_ "embed"
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
)

//go:embed templates.toml
var templatesToml string

type TemplateEntry struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	URL         string `toml:"url"`
}

type TemplateConfig struct {
	Templates []TemplateEntry `toml:"templates"`
}

func (t *TemplateConfig) Options() []huh.Option[string] {
	options := make([]huh.Option[string], len(t.Templates))
	for i, template := range t.Templates {
		options[i] = huh.NewOption(template.Name, template.URL)
	}
	return options
}

func LoadTemplates(path string) (*TemplateConfig, error) {
	var templates TemplateConfig

	if path == "" {
		_, err := toml.Decode(templatesToml, &templates)
		if err != nil {
			return nil, errors.New("failed to parse template config: " + err.Error())
		}
		return &templates, nil
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("failed to read template config: " + err.Error())
	}

	_, err = toml.Decode(string(file), &templates)
	if err != nil {
		return nil, errors.New("failed to parse template config: " + err.Error())
	}

	return &templates, nil
}
