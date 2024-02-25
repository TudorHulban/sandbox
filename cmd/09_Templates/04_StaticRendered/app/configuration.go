package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/TudorHulban/log"
	"gopkg.in/yaml.v3"
)

type HTMLPageTemplate struct {
	ContainingFolder string
	PageShell        string

	Head   string
	Meta   string
	Header string
	Body   string
	// Section , Aside string
	Article string
	Footer  string
}

type ConfigFile struct {
	AppFile string
	Backup  string
}

type InfoSite struct {
	FaviconImagePath string
	SiteLogoPath     string

	Host          string `yaml:"host"`
	ListeningPort int    `yaml:"port"`
}

type InfoArticles struct {
	ArticlesRAWFolder      string
	ArticlesRenderToFolder string
}

type ConfigurationApp struct {
	HTMLPageTemplate

	L *log.Logger `yaml:"-"`

	ConfigFile `yaml:"config-file"`

	InfoSite     `yaml:"site"`
	InfoArticles `yaml:"articles"`
}

func NewAppConfiguration(importPath string, logLevel int) (*ConfigurationApp, error) {
	if len(importPath) == 0 {
		return defaultConfiguration()
	}

	data, errRead := os.ReadFile(importPath)
	if errRead != nil {
		return nil,
			fmt.Errorf(
				"issues when loading blog articles in file %s:%w",
				importPath,
				errRead,
			)
	}

	var result struct {
		HTMLPageTemplate
		InfoSite
	}

	if errUnmarshal := json.Unmarshal(data, &result); errUnmarshal != nil {
		return nil,
			fmt.Errorf(
				"issues when unmarshaling configuration data:%w",
				errUnmarshal,
			)
	}

	return &ConfigurationApp{
			InfoSite:         result.InfoSite,
			HTMLPageTemplate: result.HTMLPageTemplate,

			L: log.NewLogger(log.DEBUG, os.Stdout, true),
		},
		nil
}

func (cfg *ConfigurationApp) SaveConfiguration(w io.Writer) (int, error) {
	configuration, errMarshal := yaml.Marshal(cfg)
	if errMarshal != nil {
		return 0,
			fmt.Errorf(
				"unmarshaling configuration: %w",
				errMarshal,
			)
	}

	return w.Write(configuration)
}

func (cfg *ConfigurationApp) SaveConfigurationTo(path string) error {
	configuration, errMarshal := yaml.Marshal(cfg)
	if errMarshal != nil {
		return fmt.Errorf(
			"unmarshaling configuration: %w",
			errMarshal,
		)
	}

	return os.WriteFile(
		path,
		configuration,
		0o600,
	)
}

func (cfg *ConfigurationApp) ReadFrom(path string) error {
	f, errOpen := os.Open(path)
	if errOpen != nil {
		return errOpen
	}
	defer f.Close()

	var configuration ConfigurationApp

	decoder := yaml.NewDecoder(f)
	if errdecode := decoder.Decode(&configuration); errdecode != nil {
		return errdecode
	}

	*cfg = configuration

	return nil
}
