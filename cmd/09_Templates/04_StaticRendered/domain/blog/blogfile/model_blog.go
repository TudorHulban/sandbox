package blogfile

import (
	"encoding/json"
	"io/ioutil"

	"os"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/cmd/config"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/blog"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/page"
	"github.com/TudorHulban/log"

	"github.com/pkg/errors"
)

var _ blog.IBlog = (*Blog)(nil)

type Blog struct {
	Data                   []article.Article
	articlesRenderToFolder string
	l                      *log.Logger
}

func NewBlogFromArticles(cfg *config.AppConfiguration, art ...article.Article) (blog.IBlog, error) {
	if len(art) == 0 {
		return nil, errors.New("no articles provided")
	}

	if cfg.ArticlesRenderToFolder == "" {
		return nil, errors.New("folder where to render articles missing")
	}

	result := Blog{
		Data:                   []article.Article{},
		articlesRenderToFolder: cfg.ArticlesRenderToFolder,
		l:                      cfg.L,
	}

	cfg.L.Debug("ArticlesRenderToFolder:", result.articlesRenderToFolder)

	for _, a := range art {
		if err := result.addArticle(a); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// NewBlog Constructor, takes a list of file names and imports them.
// The file names should point to JSON files containing article data.
func NewBlogFromFiles(cfg *config.AppConfiguration, importFiles ...string) (blog.IBlog, error) {
	if len(importFiles) == 0 {
		return nil, errors.New("no import files provided")
	}

	if cfg.ArticlesRenderToFolder == "" {
		return nil, errors.New("folder where to render articles missing")
	}

	result := Blog{
		Data:                   []article.Article{},
		articlesRenderToFolder: cfg.ArticlesRenderToFolder,
		l:                      cfg.L,
	}

	for _, f := range importFiles {
		a, errLoad := result.loadArticle(f)
		if errLoad != nil {
			return nil, errors.WithMessagef(errLoad, "could not load article from %s", f)
		}

		result.Data = append(result.Data, *a)
	}

	if len(result.Data) == 0 {
		return nil, errors.New("no articles to batch import")
	}

	return &result, nil
}

// loadArticle Loads article from file, performs article validation.
func (b *Blog) loadArticle(loadFrom string) (*article.Article, error) {
	data, errRead := os.ReadFile(loadFrom)
	if errRead != nil {
		return nil, errors.WithMessagef(errRead, "issues when loading blog article in file %s", loadFrom)
	}

	var result article.Article

	errUnmar := json.Unmarshal(data, &result)
	if errUnmar != nil {
		return nil, errors.WithMessage(errUnmar, "issues when unmarshaling blog article data")
	}

	if errValid := result.ValidateArticle(); errValid != nil {
		return nil, errValid
	}

	return &result, nil
}

// AddArticle Method to be used when adding articles as it offers input validation.
func (b *Blog) addArticle(a article.Article) error {
	if errValid := article.Article.ValidateArticle(a); errValid != nil {
		return errors.WithMessagef(errValid, "could not validate article %v", a)
	}

	b.Data = append(b.Data, a)
	return nil
}

// GetArticles Method satisfies articles interface.
func (b *Blog) GetArticles() ([]article.Article, error) {
	return b.Data, nil
}

// GetArticle Method satisfies articles interface.
func (b *Blog) GetArticle(code string) (*article.Article, error) {
	for _, v := range b.Data {
		if code == v.CODE {
			return &v, nil
		}
	}

	return nil, errors.WithMessage(nil, "no articles found")
}

// RenderArticles Main method of blog. Part of exposed interface.
func (b *Blog) RenderArticles() error {
	p, errNew := page.NewPage(page.PageArticle())
	if errNew != nil {
		return errors.WithMessage(errNew, "error creating article page")
	}

	for _, art := range b.Data {
		if errRender := p.RenderArticle(art, b.articlesRenderToFolder); errRender != nil {
			return errors.WithMessagef(errRender, "error rendering article %s", art.CODE)
		}
	}

	return nil
}

func (b *Blog) saveBlogArticles() error {
	for _, a := range b.Data {
		if err := b.saveArticle(a); err != nil {
			return err
		}
	}

	return nil
}

func (b *Blog) saveArticle(a article.Article) error {
	byteArticle, errMar := json.MarshalIndent(a, "", " ")
	if errMar != nil {
		return errMar
	}

	return ioutil.WriteFile(a.SaveToFile, byteArticle, 0644)
}
