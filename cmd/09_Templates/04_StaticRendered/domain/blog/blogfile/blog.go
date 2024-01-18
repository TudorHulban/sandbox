package blogfile

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/blog"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/page"

	"github.com/pkg/errors"
)

var _ blog.IBlog = (*Blog)(nil)

type Blog struct {
	Articles           *article.Articles
	articlesRenderPath string
}

func NewBlog(renderPath string, articles ...*article.Article) (*Blog, error) {
	if len(articles) == 0 {
		return nil,
			errors.New("no articles provided")
	}

	if len(renderPath) == 0 {
		return nil,
			errors.New("folder where to render articles missing")
	}

	articlesBlog := article.Articles(articles)

	if errValidateArticles := articlesBlog.Validate(); errValidateArticles != nil {
		return nil, errValidateArticles
	}

	return &Blog{
			Articles:           &articlesBlog,
			articlesRenderPath: renderPath,
		},
		nil
}

func (b *Blog) AddVerifiedArticle(art *article.Article) {
	(*b).Articles = append((*b).Articles, art)
}

// GetArticle Method satisfies articles interface.
func (b *Blog) GetArticle(code string) (*article.Article, error) {
	for _, v := range b.Articles {
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

	for _, art := range b.Articles {
		if errRender := p.RenderArticle(art, b.articlesRenderPath); errRender != nil {
			return errors.WithMessagef(errRender, "error rendering article %s", art.CODE)
		}
	}

	return nil
}

func (b *Blog) saveBlogArticles() error {
	for _, a := range b.Articles {
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

	return ioutil.WriteFile(a.SaveToFilePath, byteArticle, 0644)
}
