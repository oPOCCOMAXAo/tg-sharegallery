package texts

import (
	"regexp"
	"strings"
)

type replaceExpr struct {
	SrcRe     *regexp.Regexp
	SrcString string
	Replaced  string
}

type Replacer struct {
	Replaces []replaceExpr
}

func NewReplacer() *Replacer {
	return &Replacer{}
}

func (r *Replacer) AddRegexp(src *regexp.Regexp, replaced string) {
	r.Replaces = append(r.Replaces, replaceExpr{SrcRe: src, Replaced: replaced})
}

func (r *Replacer) AddString(src string, replaced string) {
	r.Replaces = append(r.Replaces, replaceExpr{SrcString: src, Replaced: replaced})
}

func (r *Replacer) Execute(value string) string {
	for _, expr := range r.Replaces {
		if expr.SrcRe != nil {
			value = expr.SrcRe.ReplaceAllString(value, expr.Replaced)
		} else {
			value = strings.ReplaceAll(value, expr.SrcString, expr.Replaced)
		}
	}

	return value
}
