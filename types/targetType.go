package types

type Target int

const (
	TARGET_SELF Target = iota
	TARGET_BLANK
)

func (t Target) String() string {
	return [...]string{
		"_self",
		"_blank",
	}[t]
}
