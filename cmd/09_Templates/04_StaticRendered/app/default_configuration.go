package app

import (
	"os"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

type Option func(cfg *ConfigurationApp) error

func defaultConfiguration(options ...Option) (*ConfigurationApp, error) {
	executableFolder, err := os.Getwd()
	if err != nil {
		return nil,
			errors.WithMessage(err, "issues when creating default configuration")
	}

	var renderToFolder string

	if _defaultArticlesRenderedFolder[:1] != "/" {
		renderToFolder = executableFolder + "/" + _defaultArticlesRenderedFolder
	} else {
		renderToFolder = executableFolder + _defaultArticlesRenderedFolder
	}

	result := &ConfigurationApp{
		SiteInfo: SiteInfo{
			ListeningPort:          _defaultListeningPort,
			ArticlesRAWFolder:      _defaultArticlesRAWFolder,
			ArticlesRenderToFolder: renderToFolder,
		},

		AppConfigFile: _defaultAppConfigurationFileName,

		HTMLPageTemplate: HTMLPageTemplate{
			ContainingFolder: ".." + executableFolder + "/static/assets",
			PageShell:        "01_page_shell.gohtml",
			Head:             "02_head.gohtml",
			Meta:             "03_meta.gohtml",
			Header:           "04_header.gohtml",
			Body:             "05_body.gohtml",
			Article:          "06_article.gohtml",
			Footer:           "07_footer.gohtml",
		},

		L: log.NewLogger(log.DEBUG, os.Stdout, true),
	}

	// moved below initialization in order to use the logger
	result.L.Print(renderToFolder)

	if _, err := os.Stat(renderToFolder); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(renderToFolder, os.ModePerm)
		}
	}

	result.SaveConfiguration(os.Stdout) // TODO: save to file

	return result, nil
}
