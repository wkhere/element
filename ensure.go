package element

import "regexp"

func EnsureText(container ElementData) string {
	return Extract{container}.GetT()
}

func EnsureAttr(container ElementData, attr string) string {
	return Extract{container}.GetA(attr)
}

func EnsureVisible(container ElementData) bool {
	return Extract{container}.IsVisible()
}

func EnsureMatch(s string, re *regexp.Regexp) bool {
	return re.MatchString(s)
}
