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

type App struct {
	ConfigurationApp
	AppServices

	Templates         map[TemplateName]TemplateContents
	Blog              *blog.Blog
	ECommerceProducts []products.Product
}

type TemplateName string
type TemplateContents string

func NewApp(configurationFilePath string) (*App, error) {
	return &App{},
		nil
}

func (a *App) Start() error {
	return nil
}
