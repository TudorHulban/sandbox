package page

func PageArticle() Option {
	return func(p *Page) error {
		// composing article page
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

		p.Add(p.GetCurrentPos()+1, &Node{
			Name: "Body Start",
			HTML: BODYStart,
		})

		p.Add(p.GetCurrentPos()+1, &Node{
			Name: "Footer",
			HTML: FOOTER,
		})

		p.Add(p.GetCurrentPos()+1, &Node{
			Name: "Page End",
			HTML: PageEnd,
		})

		return nil
	}
}
