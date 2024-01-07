package types

type Target int

const (
	SELF Target = iota
	BLANK
)

func (t Target) String() string {
	return [...]string{
		"_self",
		"_blank",
	}[t]
}
