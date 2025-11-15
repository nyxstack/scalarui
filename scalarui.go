package scalarui

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"
)

// Embed the HTML template
//
//go:embed template.html
var htmlTemplate string

// TemplateData represents the data passed to the HTML template
type TemplateData struct {
	Title        string
	Description  string
	Favicon      string
	CustomCSS    string
	Variables    map[string]string
	ConfigJSON   template.JS
	HotReloadURL string
}

// ScalarUI represents a configured Scalar UI instance
type ScalarUI struct {
	config *Config
}

// New creates a new ScalarUI instance with the given configuration
func New(config *Config) *ScalarUI {
	if config == nil {
		config = NewConfig()
	}
	return &ScalarUI{
		config: config,
	}
}

// NewWithDefaults creates a new ScalarUI instance with default configuration
func NewWithDefaults() *ScalarUI {
	return New(NewConfig())
}

// SetConfig updates the configuration
func (s *ScalarUI) SetConfig(config *Config) {
	s.config = config
}

// GetConfig returns the current configuration
func (s *ScalarUI) GetConfig() *Config {
	return s.config
}

// Render generates the HTML string with the configured options
func (s *ScalarUI) Render() (string, error) {
	return renderTemplate(s.config)
}

// renderTemplate renders the HTML template with the given configuration
func renderTemplate(config *Config) (string, error) {
	// Convert config to JSON for JavaScript
	configBytes, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return "", err
	}

	// Prepare template data
	data := TemplateData{
		Title:        config.Title,
		Description:  config.Description,
		Favicon:      config.Favicon,
		CustomCSS:    config.CustomCSS,
		Variables:    config.Variables,
		ConfigJSON:   template.JS(configBytes),
		HotReloadURL: config.HotReloadURL,
	}

	// Parse and execute template
	tmpl, err := template.New("scalar").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
