package types

type Page struct {
	Title       string
	Description string
	Index       string
	Follow      string
	Routes      []Route
	Content     []Content[Block]
}

func NewPage() Page {
	return Page{
		Title:       "",
		Description: "",
		Index:       "index",
		Follow:      "follow",
		Routes:      []Route{},
		Content:     []Content[Block]{},
	}
}

func (p *Page) SetTitle(title string) {
	p.Title = title
}

func (p *Page) SetDescription(description string) {
	p.Description = description
}

func (p *Page) SetIndex(index string) {
	p.Index = index
}

func (p *Page) SetFollow(follow string) {
	p.Follow = follow
}

func (p *Page) SetRoutes(routes []Route) {
	p.Routes = routes
}

func (p *Page) SetContent(content []Content[Block]) {
	p.Content = content
}

func (p *Page) GetTitle() string {
	return p.Title
}

func (p *Page) GetDescription() string {
	return p.Description
}

func (p *Page) GetIndex() string {
	return p.Index
}

func (p *Page) GetFollow() string {
	return p.Follow
}

func (p *Page) GetRoutes() []Route {
	return p.Routes
}

func (p *Page) GetContent() []Content[Block] {
	return p.Content
}

func (p *Page) AppendRoute(route Route) {
	p.Routes = append(p.Routes, route)
}

func (p *Page) AppendContent(content Content[Block]) {
	p.Content = append(p.Content, content)
}
