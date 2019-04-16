package element

type elementData interface {
	Text() (string, error)
	Attribute(string) (string, error)
	Enabled() (bool, error)
	Visible() (bool, error)
}

type Extract struct{ elementData }

func (e Extract) GetT() string {
	text, err := e.Text()
	if err != nil {
		Die("extract text error: %v", err)
	}
	return text
}

func (e Extract) GetV() string {
	v, err := e.Attribute("value")
	if err != nil {
		Die("extract value error: %v", err)
	}
	return v
}

func (e Extract) GetA(attr string) string {
	v, err := e.Attribute(attr)
	if err != nil {
		Die("extract attr error:: %v", err)
	}
	return v
}

func (e Extract) IsVisible() bool {
	ok, err := e.Visible()
	if err != nil {
		Die("extract isVisible error: %v", err)
	}
	return ok
}
