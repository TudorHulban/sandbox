package servicerender

import (
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/page"
	"github.com/pkg/errors"
)

type Service struct {
	optionArticle page.Option
}

func (s *Service) RenderArticle(article *article.Article, toFolder string) error {
	pag, errNew := page.NewPage(s.optionArticle)
	if errNew != nil {
		return errors.WithMessage(errNew, "error creating article page")
	}

	return pag.RenderArticle(article, toFolder)
}

func (s *Service) RenderArticles(articles *article.Articles, toFolder string) error {
	for _, art := range *articles {
		if errRender := s.RenderArticle(art, toFolder); errRender != nil {
			return errRender
		}
	}

	return nil
}
