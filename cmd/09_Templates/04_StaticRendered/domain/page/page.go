package page

import (
	"errors"
	"html/template"
	"os"
	"strings"

	"github.com/TudorHulban/GolangSandbox/cmd/09_Templates/04_StaticRendered/domain/article"
)

type SiteInfo struct {
	Title    string
	Subtitle string
}

type Node struct {
	Name string
	HTML string
}

type Page struct {
	Nodes [][]*Node
}

type Option func(p *Page) error

func NewPage(options ...Option) (*Page, error) {
	result := &Page{
		Nodes: [][]*Node{},
	}

	for _, option := range options {
		if errOptionApply := option(result); errOptionApply != nil {
			return nil, errOptionApply
		}
	}

	return result, nil
}

func (p *Page) Add(pos uint, n *Node) {
	if pos <= uint(len(p.Nodes)) {
		p.Nodes[pos-1] = append(p.Nodes[pos-1], n)

		return
	}

	p.Nodes = append(p.Nodes, []*Node{n})
}

func (p *Page) GetCurrentPos() uint {
	return uint(len(p.Nodes))
}

func (p *Page) GetHTML() []string {
	var result []string

	for _, nodeLevel := range p.Nodes {
		for _, node := range nodeLevel {
			result = append(result, node.HTML)
		}
	}

	return result
}

// GetHTML Method returns page HTML as string.
func (p *Page) GetString() string {
	return strings.Join(p.GetHTML(), "\n")
}

func (p *Page) RenderArticle(a article.Article, renderToFolder string) error {
	renderToFolder = strings.Trim(renderToFolder, " ")

	if len(renderToFolder) == 0 {
		return errors.New("folder where to render articles missing")
	}

	t, errParse := template.New("").Parse(p.GetString())
	if errParse != nil {
		return errParse
	}

	toPath := renderToFolder + "/" + strings.ToLower(a.CODE) + ".html"

	f, errCreate := os.Create(toPath)
	if errCreate != nil {
		return errCreate
	}
	defer f.Close()

	if errExec := t.Execute(f, a); errExec != nil {
		return errExec
	}

	return nil
}
