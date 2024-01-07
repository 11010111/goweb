package types

type Confirmation struct {
	Headline  string
	Bodytext  string
	Firstname string
	Lastname  string
	Email     string
}

func NewConfirmation() Confirmation {
	return Confirmation{}
}

func (c *Confirmation) SetHeadline(headline string) {
	c.Headline = headline
}

func (c *Confirmation) SetBodytext(bodytext string) {
	c.Bodytext = bodytext
}

func (c *Confirmation) SetFirstname(firstname string) {
	c.Firstname = firstname
}

func (c *Confirmation) SetLastname(lastname string) {
	c.Lastname = lastname
}

func (c *Confirmation) SetEmail(email string) {
	c.Email = email
}

func (c *Confirmation) GetHeadline() string {
	return c.Headline
}

func (c *Confirmation) GetBodytext() string {
	return c.Bodytext
}

func (c *Confirmation) GetFirstname() string {
	return c.Firstname
}

func (c *Confirmation) GetLastname() string {
	return c.Lastname
}

func (c *Confirmation) GetEmail() string {
	return c.Email
}
