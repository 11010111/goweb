package types

type Route struct {
	Title    string
	Href     string
	Target   string
	Relation string
}

func NewRoute() Route {
	return Route{
		Title:    "Link",
		Href:     "/",
		Target:   "_self",
		Relation: "",
	}
}

func (r *Route) SetTitle(title string) {
	r.Title = title
}

func (r *Route) SetHref(href string) {
	r.Href = href
}

func (r *Route) SetTarget(target string) {
	r.Target = target
}

func (r *Route) SetRelation(relation string) {
	r.Relation = relation
}

func (r *Route) GetTitle() string {
	return r.Title
}

func (r *Route) GetHref() string {
	return r.Href
}

func (r *Route) GetTarget() string {
	return r.Target
}

func (r *Route) GetRelation() string {
	return r.Relation
}
