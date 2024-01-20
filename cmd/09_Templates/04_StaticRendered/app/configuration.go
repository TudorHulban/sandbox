package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
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

type SiteInfo struct {
	ListeningPort    int    `yaml:"port"`
	Host             string `yaml:"host"`
	FaviconImagePath string
	SiteLogoPath     string

	ArticlesRAWFolder      string
	ArticlesRenderToFolder string
}

type ConfigurationApp struct {
	AppConfigFile    string
	SaveToConfigFile string

	L *log.Logger `json:"-"`

	SiteInfo `yaml:"site"`
	HTMLPageTemplate
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
		SiteInfo
		HTMLPageTemplate
	}

	if errUnmarshal := json.Unmarshal(data, &result); errUnmarshal != nil {
		return nil,
			fmt.Errorf(
				"issues when unmarshaling configuration data:%w",
				errUnmarshal,
			)
	}

	return &ConfigurationApp{
			SiteInfo:         result.SiteInfo,
			HTMLPageTemplate: result.HTMLPageTemplate,
			L:                log.NewLogger(log.DEBUG, os.Stdout, true),
		},
		nil
}

func (cfg *ConfigurationApp) SaveConfiguration(w io.Writer) (n int, err error) {
	// configuration, errMarshal := json.MarshalIndent(cfg, "", " ")
	configuration, errMarshal := yaml.Marshal(cfg)
	if errMarshal != nil {
		return 0,
			errors.WithMessage(errMarshal, "could not unmarshal configuration")
	}

	return w.Write(configuration)
}

func (cfg *ConfigurationApp) SaveConfigurationTo(path string) error {
	configuration, errMarshal := yaml.Marshal(cfg)
	if errMarshal != nil {
		return errors.WithMessage(errMarshal, "could not unmarshal configuration")
	}

	return os.WriteFile(
		path,
		configuration,
		0644,
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
