package main

import (
	"regexp"
	"strings"
)

var (
	// used to replace multiple non-alphanumerics with a space to sanitize names
	invalidNameCharReplacement  = regexp.MustCompile("[^a-zA-Z0-9]+")
	invalidShortNameReplacement = regexp.MustCompile("[^a-z0-9_]+")
)

// sanitizedName uses the lowercase name field or replaces
// underscores with spaces in the short code
func sanitizedName(e EmojiInfo) string {
	name := e.Name
	if len(name) == 0 {
		name = e.ShortName
	}
	name = strings.ToLower(name)
	return invalidNameCharReplacement.ReplaceAllString(name, " ")
}

// sanitizedShortName uses the lowercase short name field or replaces
// invalid characters in the name with an underscore.
func sanitizedShortName(e EmojiInfo) string {
	name := e.ShortName
	if len(name) == 0 {
		name = e.Name
	}
	name = strings.ToLower(name)
	return invalidShortNameReplacement.ReplaceAllString(name, "_")
}

func uniqueKeywords(e EmojiInfo) (keywords []string) {
	added := map[string]struct{}{}
	for _, phrase := range append([]string{e.Name, e.ShortName}, append(e.ShortNames, e.Texts...)...) {
		if _, exists := added[phrase]; exists {
			continue
		}
		keywords = append(keywords, phrase)
		added[phrase] = struct{}{}
	}
	return keywords
}
