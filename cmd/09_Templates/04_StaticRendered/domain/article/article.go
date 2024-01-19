package article

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Article struct {
	CODE    string `json:"code"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Content string `json:"content"`

	HTMLTemplateFile  string `json:"html"`
	FeaturedImagePath string `json:"featuredimage"`

	Created     uint64 `json:"created"` // UNIX time seconds
	LastUpdated uint64 `json:"updated"` // UNIX time seconds
	IsVisible   bool   `json:"visible"`

	RelatedProductsSKUs       []uint64 `json:"relatedsku"`
	RelatedProductsCategories []string `json:"relatedcateg"`
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

func (a *Article) Validate() error {
	return nil
}

func (a *Article) SaveTo(w io.Writer) (int, error) {
	byteArticle, errMarshal := json.MarshalIndent(a, "", " ")
	if errMarshal != nil {
		return 0,
			errMarshal
	}

	return w.Write(byteArticle)
}

func (a *Article) savePath(folder string) string {
	if strings.LastIndex(folder, _slash) == -1 {
		return folder + "/" + a.CODE + _articleFileExtension
	}

	return folder + a.CODE + _articleFileExtension
}

func (a *Article) SaveToPath(folder string) (int, error) {
	f, errCreate := os.Create(
		a.savePath(folder),
	)
	if errCreate != nil {
		return 0,
			errCreate
	}

	return a.SaveTo(f)
}
