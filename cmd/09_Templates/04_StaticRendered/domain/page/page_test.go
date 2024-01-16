package page

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	p, _ := NewPage()

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Page Start",
		HTML: PageStart,
	})

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Head",
		HTML: HEAD,
	})

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Body Start",
		HTML: BODYStart,
	})

	level := p.GetCurrentPos() + 1
	p.Add(level, &Node{
		Name: "2a",
		HTML: "x",
	})

	p.Add(level, &Node{
		Name: "2b",
		HTML: "y",
	})

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Body End",
		HTML: BODYEnd,
	})

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Footer",
		HTML: FOOTER,
	})

	p.Add(p.GetCurrentPos()+1, &Node{
		Name: "Page End",
		HTML: PageEnd,
	})

	//require.Equal(t, len(p.Nodes), 4)

	fmt.Println("Page:")
	fmt.Println(p.GetString())
}
