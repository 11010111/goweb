package main

type Route struct {
	Title    string
	Href     string
	Target   string
	Relation string
}

type Navigation struct {
	Routes []Route
}

type Hero struct {
	Headline string
	Bodytext string
}

type Text struct {
	Headline string
	Bodytext string
}

type Form struct {
	Headline string
}

type Confirmation struct {
	Headline  string
	Bodytext  string
	Firstname string
	Lastname  string
	Email     string
}

type Block[T any] struct {
	Type    string
	Content T
}

type Page struct {
	Title       string
	Description string
	Index       string
	Follow      string
	Navigation  Navigation
	Body        []Block[any]
}
