package types

type Hero struct {
	Headline string
	Bodytext string
}

func NewHero() Hero {
	return Hero{}
}

func (h *Hero) SetHeadline(headline string) {
	h.Headline = headline
}

func (h *Hero) SetBodytext(bodytext string) {
	h.Bodytext = bodytext
}

func (h *Hero) GetHeadline() string {
	return h.Headline
}

func (h *Hero) GetBodytext() string {
	return h.Bodytext
}
