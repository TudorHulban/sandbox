package app

import (
	"fmt"
	"os"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/app/constants"
	"github.com/TudorHulban/log"
)

type Option func(cfg *ConfigurationApp) error

func defaultConfiguration(options ...Option) (*ConfigurationApp, error) {
	executableFolder, errGetWorkingDirectory := os.Getwd()
	if errGetWorkingDirectory != nil {
		return nil,
			fmt.Errorf(
				"creating default configuration: %w",
				errGetWorkingDirectory,
			)
	}

	var renderToFolder string

	if _defaultArticlesRenderedFolder[:1] != "/" {
		renderToFolder = executableFolder + "/" + _defaultArticlesRenderedFolder
	} else {
		renderToFolder = executableFolder + _defaultArticlesRenderedFolder
	}

	result := &ConfigurationApp{
		InfoSite: InfoSite{
			ListeningPort: constants.DefaultListeningPort,
		},

		InfoArticles: InfoArticles{
			ArticlesRAWFolder:      _defaultArticlesRAWFolder,
			ArticlesRenderToFolder: renderToFolder,
		},

		ConfigFile: ConfigFile{
			AppFile: constants.ConfigFilePath,
		},

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

	if _, err := os.Stat(renderToFolder); err != nil {
		if os.IsNotExist(err) {
			_ = os.Mkdir(renderToFolder, os.ModePerm)
		}
	}

	result.SaveConfigurationTo(result.AppFile)

	return result, nil
}
