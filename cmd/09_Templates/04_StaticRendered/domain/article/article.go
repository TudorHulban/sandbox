package article

import (
	"encoding/json"
	"fmt"
	"os"
)

// Article represents blog entry to help conversion and is common to all blog implementations.
type Article struct {
	SaveToFilePath string `json:"file"`
	CODE           string `json:"code"`
	Name           string `json:"name"`
	Author         string `json:"author"`
	Content        string `json:"content"`

	HTMLTemplateFile  string `json:"html"`
	FeaturedImagePath string `json:"featuredimage"`

	IsVisible   bool   `json:"visible"`
	Created     uint64 `json:"created"` // UNIX time seconds
	LastUpdated uint64 `json:"updated"` // UNIX time seconds

	RelatedProductsSKUs       []uint64 `json:"relatedsku"`
	RelatedProductsCategories []string `json:"relatedcateg"`
}

func (a *Article) Validate() error {
	return nil
}

func NewArticle(fromPath string) (*Article, error) {
	data, errRead := os.ReadFile(fromPath)
	if errRead != nil {
		return nil,
			fmt.Errorf(
				"reading file %s: %w",
				fromPath,
				errRead,
			)
	}

	var result Article

	errUnmarshal := json.Unmarshal(data, &result)
	if errUnmarshal != nil {
		return nil,
			fmt.Errorf(
				"unmarshaling article from file %s: %w",
				fromPath,
				errUnmarshal,
			)
	}

	if errValid := result.Validate(); errValid != nil {
		return nil,
			errValid
	}

	return &result, nil
}
