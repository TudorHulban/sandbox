package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
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
	ListeningPort    int
	FaviconImagePath string
	SiteLogoPath     string

	ArticlesRAWFolder      string
	ArticlesRenderToFolder string
}

type ConfigurationApp struct {
	AppConfigFile    string
	SaveToConfigFile string
	L                *log.Logger

	SiteInfo
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
	configuration, errMarshal := json.MarshalIndent(cfg, "", " ")
	if errMarshal != nil {
		return 0,
			errors.WithMessage(errMarshal, "could not unmarshal configuration")
	}

	return w.Write(configuration)
}
