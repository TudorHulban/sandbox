package blog

import (
	"testing"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
	"github.com/stretchr/testify/require"
)

func TestNewBlog(t *testing.T) {
	blog, errNew := NewBlog(article.DefaultArticles()...)
	require.NoError(t, errNew)
	require.NotNil(t, blog)
}
