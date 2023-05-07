package config

import (
	"github.com/TudorHulban/GoTemplating/internal/article"
	products "github.com/TudorHulban/GoTemplating/internal/ecommerce"
	"github.com/TudorHulban/log"
)

type SiteInfo struct {
	ListeningPort    int
	FaviconImagePath string
	SiteLogoPath     string

	ArticlesRAWFolder      string
	ArticlesRenderToFolder string
}

// HTMLPageTemplates Consolidates HTML page templates.
// All templates shold share same containing folder and fields should be file names only.
type HTMLPageTemplates struct {
	ContainingFolder string
	PageShell        string
	Head             string
	Meta             string
	Header           string
	Body             string
	// Section , Aside string
	Article string
	Footer  string
}

// AppConfiguration Structure holding application configuration.
type AppConfiguration struct {
	AppConfigFile    string
	SaveToConfigFile string
	L                *log.Logger

	SiteInfo
	HTMLPageTemplates
}

type App struct {
	AppConfiguration

	Templates         map[TemplateName]TemplateContents
	BlogArticles      []article.Article
	ECommerceProducts []products.Product
}
