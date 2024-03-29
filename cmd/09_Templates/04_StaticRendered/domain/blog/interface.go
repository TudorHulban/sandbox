package blog

// IArticles interface added for different implementations / types of persistence.
// Blog types should save automatically articles to persistence when there is a change.
type IBlog interface {
	RenderArticles() error

	//RenderArticlesSummarization() error

	//GetArticles() ([]Article, error)
	//GetArticle(code string) (*Article, error)

	//GetRelatedArticles(sku uint64, howMany uint8) (Article, error)
	//GetCategoryArticles(category string, howMany uint8) (Article, error)
}
