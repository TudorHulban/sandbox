package blogfile

import (
	"testing"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/stretchr/testify/require"
)

// test load article
// test load articles

// test save article
// test save articles

func TestBlogArticles(t *testing.T) {
	b, err := NewBlog(article.DefaultArticles()...)
	require.Nil(t, err)

	require.Nil(t, b.(*Blog).saveBlogArticles())
	require.Nil(t, b.RenderArticles())
}

func TestBlogFiles(t *testing.T) {
	files := make([]string, len(article.DefaultArticles()))

	for i, art := range article.DefaultArticles() {
		files[i] = art.SaveToFilePath
	}

	_, err := NewBlogFromFiles(files...)
	require.Nil(t, err)
}
