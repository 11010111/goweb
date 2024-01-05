package main

type Relation int

const (
	Default Relation = iota
	Noopener
	Noreferrer
	Nofollow
	Preload
	Preconnect
	Stylesheet
)

func (rel Relation) String() string {
	return [...]string{
		"",
		"noopener",
		"noreferrer",
		"nofollow",
		"preload",
		"preconnect",
		"stylesheet",
	}[rel]
}
