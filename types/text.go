package types

type Text struct {
	Headline string
	Bodytext string
}

func NewText() Text {
	return Text{}
}

func (t *Text) SetHeadline(headline string) {
	t.Headline = headline
}

func (t *Text) SetBodytext(bodytext string) {
	t.Bodytext = bodytext
}

func (t *Text) GetHeadline() string {
	return t.Headline
}

func (t *Text) GetBodytext() string {
	return t.Bodytext
}
