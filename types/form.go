package types

type Form struct {
	Headline string
	Message  string
}

func NewForm() Form {
	return Form{}
}

func (f *Form) SetHeadline(headline string) {
	f.Headline = headline
}

func (f *Form) SetMessage(message string) {
	f.Message = message
}

func (f *Form) GetHeadline() string {
	return f.Headline
}

func (f *Form) GetMessage() string {
	return f.Message
}
