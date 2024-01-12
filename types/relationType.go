package types

type Relation int

const (
	REL_NONE Relation = iota
	REL_NOOPENER
	REL_NOREFERRER
	REL_NOFOLLOW
	REL_PRELOAD
	REL_PRECONNECT
	REL_STYLESHEET
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
