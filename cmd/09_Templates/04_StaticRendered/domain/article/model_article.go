package article

// Article represents blog entry to help conversion and is common to all blog implementations.
type Article struct {
	SaveToFile        string `json:"file"` // particular to blog file
	CODE              string `json:"code"`
	Name              string `json:"name"`
	Author            string `json:"author"`
	Content           string `json:"content"`
	HTMLTemplateFile  string `json:"html"`
	FeaturedImagePath string `json:"featuredimage"`

	IsVisible   bool   `json:"visible"`
	Created     uint64 `json:"created"` // UNIX time seconds
	LastUpdated uint64 `json:"updated"` // UNIX time seconds

	RelatedProductsSKUs       []uint64 `json:"relatedsku"`
	RelatedProductsCategories []string `json:"relatedcateg"`
}

func (a Article) ValidateArticle() error {
	return nil
}
