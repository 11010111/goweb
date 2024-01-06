package main

type Relation int

const (
	NONE Relation = iota
	NOOPENER
	NOREFERRER
	NOFOLLOW
	PRELOAD
	PRECONNECT
	STYLESHEET
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
