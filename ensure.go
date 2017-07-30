package element

import "regexp"

func EnsureText(container elementData) string {
	return Extract{container}.GetT()
}

func EnsureAttr(container elementData, attr string) string {
	return Extract{container}.GetA(attr)
}

func EnsureVisible(container elementData) bool {
	return Extract{container}.IsVisible()
}

func EnsureMatch(s string, re *regexp.Regexp) bool {
	return re.MatchString(s)
}
