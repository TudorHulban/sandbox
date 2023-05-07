package blogfile

import (
	"os"
	"testing"

	"github.com/TudorHulban/GoTemplating/internal/article"
	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/require"
)

// test load article
// test load articles

// test save article
// test save articles

var l = log.NewLogger(log.DEBUG, os.Stdout, true)

func TestBlogArticles(t *testing.T) {
	b, err := NewBlogFromArticles(l, article.DefaultArticles()...)
	require.Nil(t, err)

	require.Nil(t, b.(*Blog).saveBlogArticles())
	require.Nil(t, b.RenderArticles())
}

func TestBlogFiles(t *testing.T) {
	files := make([]string, len(article.DefaultArticles()))

	for i, art := range article.DefaultArticles() {
		files[i] = art.SaveToFile
	}

	_, err := NewBlogFromFiles(l, files...)
	require.Nil(t, err)
}
