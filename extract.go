package element

type ElementData interface {
	Text() (string, error)
	Attribute(string) (string, error)
	Count() (int, error)
	Enabled() (bool, error)
	Visible() (bool, error)
}

type Extract struct{ ElementData }

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

func (e Extract) GetCount() int {
	n, err := e.Count()
	if err != nil {
		die("extract count error: %s", err)
	}
	return n
}

func (e Extract) IsEnabled() bool {
	ok, err := e.Enabled()
	if err != nil {
		Die("extract isEnabled error: %v", err)
	}
	return ok
}

func (e Extract) IsVisible() bool {
	ok, err := e.Visible()
	if err != nil {
		Die("extract isVisible error: %v", err)
	}
	return ok
}
