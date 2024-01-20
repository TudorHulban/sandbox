package app

import (
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/blog"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/products"
	servicearticle "github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/services/service-article"
	servicerender "github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/services/service-render"
)

type AppServices struct {
	ServiceArticle *servicearticle.Service

	ServiceRender *servicerender.Service
}

type TemplateName string
type TemplateContents string

type App struct {
	ConfigurationApp
	AppServices

	Templates         map[TemplateName]TemplateContents
	Blog              *blog.Blog
	ECommerceProducts []products.Product
}

func NewApp(configurationFilePath string) (*App, error) {
	configurationDefault, errConfig := defaultConfiguration()
	if errConfig != nil {
		return nil,
			errConfig
	}

	app := App{
		ConfigurationApp: *configurationDefault,

		AppServices: AppServices{
			ServiceArticle: &servicearticle.Service{},
			ServiceRender:  &servicerender.Service{},
		},
	}

	articles, errLoadArticles := app.ServiceArticle.ArticlesFromFolder(app.ArticlesRAWFolder)
	if errLoadArticles != nil {
		return nil,
			errLoadArticles
	}

	blog, errNew := blog.NewBlog(*articles...)
	if errNew != nil {
		return nil,
			errNew
	}

	app.Blog = blog

	return &app,
		nil

}

func (a *App) Start() error {
	return nil
}
