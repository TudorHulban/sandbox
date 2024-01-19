package app

import (
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/blog"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/products"
)

type App struct {
	ConfigurationApp

	Templates         map[TemplateName]TemplateContents
	Blog              *blog.Blog
	ECommerceProducts []products.Product
}

type TemplateName string
type TemplateContents string

func NewApp(cfg ConfigurationApp) (*App, error) {
	return nil, nil
}
