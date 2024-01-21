package blog

import (
	"fmt"
	"sync"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"

	"github.com/pkg/errors"
)

type Blog struct {
	articles []*article.Article

	mu sync.Mutex
}

func NewBlog(articles ...*article.Article) (*Blog, error) {
	if len(articles) == 0 {
		return nil,
			errors.New("no articles provided")
	}

	articlesBlog := article.Articles(articles)

	if errValidateArticles := articlesBlog.Validate(); errValidateArticles != nil {
		return nil, errValidateArticles
	}

	return &Blog{
			articles: articlesBlog,
		},
		nil
}

func (b *Blog) AddArticle(art *article.Article) error {
	if errValidate := art.Validate(); errValidate != nil {
		return errValidate
	}

	b.mu.Lock()

	b.articles = append(b.articles, art)

	b.mu.Unlock()

	return nil
}

func (b *Blog) GetArticle(code string) (*article.Article, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, article := range b.articles {
		if code == article.CODE {
			return article, nil
		}
	}

	return nil,
		fmt.Errorf(
			"article not found for CODE: %s",
			code,
		)
}

func (b *Blog) GetNumberArticles() int {
	return len(b.articles)
}
