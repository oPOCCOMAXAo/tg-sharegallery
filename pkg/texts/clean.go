package texts

import (
	"regexp"
	"strings"
)

func GetHTMLCleaner() *Replacer {
	res := NewReplacer()

	res.AddString(`&nbsp;`, " ")
	res.AddString(`&lt;`, "<")
	res.AddString(`&gt;`, ">")
	res.AddString(`&quot;`, `"`)
	res.AddString(`&apos;`, `'`)
	res.AddString(`&amp;`, "&")
	res.AddString(`&#039;`, "'")
	res.AddString(`<br>`, "\n")
	res.AddString(`<hr>`, "\n")
	res.AddRegexp(regexp.MustCompile(`<[^<>]+>`), " ")
	res.AddRegexp(regexp.MustCompile(`[\t ]+`), " ")

	// Trim beginning and ending spaces
	res.AddRegexp(regexp.MustCompile(`(?m)^\s+|\s+$`), "")

	return res
}

//nolint:gochecknoglobals
var htmlCleaner = GetHTMLCleaner()

// CleanHTML removes all HTML tags from the given string, replacing them with valid characters.
func CleanHTML(html string) string {
	html = htmlCleaner.Execute(html)
	html = strings.TrimSpace(html)

	return html
}
