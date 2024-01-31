package app

import (
	"net/http"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/blog"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/products"
	servicearticle "github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/services/service-article"
	servicerender "github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/services/service-render"
	transportfiber "github.com/TudorHulban/GolangSandbox/infra/transport-fiber"
)

type AppServices struct {
	ServiceArticle *servicearticle.Service
	ServiceRender  *servicerender.Service
}

type TemplateName string
type TemplateContents string

type App struct {
	ConfigurationApp
	AppServices
	transport *transportfiber.Transport

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

	transport, errNewTransport := transportfiber.NewTransport(
		&transportfiber.ParamTransportSocket{
			Port: "3000",
			Host: "localhost",
		},
	)
	if errNewTransport != nil {
		return nil,
			errNewTransport
	}

	app := App{
		ConfigurationApp: *configurationDefault,

		AppServices: AppServices{
			ServiceArticle: &servicearticle.Service{},
			ServiceRender:  &servicerender.Service{},
		},

		transport: transport,
	}

	app.addRoutes()

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

func (a *App) addRoutes() {
	a.transport.AddRoute(
		&transportfiber.ParamAddRoute{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: a.HandlerAll(),
		},
	)

	a.transport.AddRoute(
		&transportfiber.ParamAddRoute{
			Method:  http.MethodGet,
			Path:    "/:code",
			Handler: a.HandlerGetArticle(),
		},
	)
}

func (a *App) Start() error {
	return a.transport.Start()
}
