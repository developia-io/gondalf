package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplates(t *testing.T) {
	// Create a temporary TOML file
	tempFile, err := os.CreateTemp("", "template_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write a sample template to the file
	_, err = tempFile.Write([]byte(`
[[templates]]
name = "Test Template"
description = "This is a test template"
url = "http://example.com"
`))
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Load the templates from the file
	templates, err := LoadTemplates(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check the loaded templates
	assert.Equal(t, 1, len(templates.Templates))
	assert.Equal(t, "Test Template", templates.Templates[0].Name)
	assert.Equal(t, "This is a test template", templates.Templates[0].Description)
	assert.Equal(t, "http://example.com", templates.Templates[0].URL)
}
