package servicearticle

import (
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/app/constants"
	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/TudorHulban/GolangSandbox/helpers"
)

type Service struct{}

func (s *Service) ArticlesFromFolder(folder string) (*article.Articles, error) {
	articleFileNames, errGet := helpers.GetFiles(folder, constants.ExtensionArticleFile)
	if errGet != nil {
		return nil,
			errGet
	}

	var result article.Articles

	for _, articleFileName := range articleFileNames {
		newArticle, errNew := article.NewArticle(articleFileName)
		if errNew != nil {
			return nil,
				errNew
		}

		result = append(result, newArticle)
	}

	return &result,
		nil
}
