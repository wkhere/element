package element

import "regexp"

func AssertEq(v, expected interface{}) {
	if expected != v {
		die("expected `%v` got `%v`", expected, v)
	}
}

func AssertTextEq(container elementData, expected string) {
	s := EnsureText(container)
	AssertEq(s, expected)
}

func AssertAttrEq(container elementData, attr string, expected string) {
	v := EnsureAttr(container, attr)
	AssertEq(v, expected)
}

func AssertMatch(v string, re *regexp.Regexp) {
	if !EnsureMatch(v, re) {
		Die("no match for pattern `%s` in `%s`", re, v)
	}
}

func AssertNotMatch(v string, re *regexp.Regexp) {
	if EnsureMatch(v, re) {
		Die("unexpected match for pattern `%s` in `%s`", re, v)
	}
}

func AssertTextMatch(container elementData, re *regexp.Regexp) {
	AssertMatch(EnsureText(container), re)
}

func AssertTextNotMatch(container elementData, re *regexp.Regexp) {
	AssertNotMatch(EnsureText(container), re)
}
