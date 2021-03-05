package emoji

import (
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// SearchIndex is allows a keyword-based search of the emoji dataset.
type SearchIndex struct {
	options searchOptionSet
	// keyword from All array
	keywordStrings []string
	// index in All of the keyword.
	keywordIndexes []int
}

// NewSearchIndex creates a keyword fuzzy search index.
func NewSearchIndex(opts ...SearchOption) *SearchIndex {
	var (
		options searchOptionSet
		// keyword from All array
		keywordStrings []string
		// index in All of the keyword.
		keywordIndexes []int
	)

	for _, optionFunc := range opts {
		optionFunc(&options)
	}
	for i, info := range All {
		for _, term := range info.AlternateNames {
			keywordStrings = append(keywordStrings, term)
			keywordIndexes = append(keywordIndexes, i)
		}
	}
	return &SearchIndex{
		options:        options,
		keywordStrings: keywordStrings,
		keywordIndexes: keywordIndexes,
	}
}

// Search performs a fuzzy search on the emoji keywords to find a matching symbol.
// The response array will contain up to the configured limit number of terms that
// are less than the maximum distance from the search term.
//
// Options provided directly to the search term override the defaults passed
// to NewSearchIndex.
func (si *SearchIndex) Search(query string, opts ...SearchOption) []Info {
	ranks := fuzzy.RankFindNormalizedFold(query, si.keywordStrings)
	sort.Sort(ranks)

	options := si.options
	for _, optionFunc := range opts {
		optionFunc(&options)
	}

	results := make([]Info, 0, options.Limit)
	for _, rank := range ranks {
		if options.MaxDistance > 0 && rank.Distance > options.MaxDistance {
			break
		}
		idx := si.keywordIndexes[rank.OriginalIndex]
		if idx < len(All) {
			results = append(results, All[idx])
		}
		if options.Limit > 0 && len(results) >= options.Limit {
			break
		}
	}
	return results
}

// searchOptionSet collects values from multiple search options.
// It is internal so consumers need to use the `WithXX(...)`
// utilities to modify an option set.
type searchOptionSet struct {
	MaxDistance int
	Limit       int
}

// SearchOption represents an option that is used to search the dataset.
type SearchOption func(option *searchOptionSet)

// WithMaxDistance configures the maximum Levenshtein distance for keyword matches.
// A 0 value means no maximum distance and the result will always contain the maximum
// limit of results.
func WithMaxDistance(maxDistance int) SearchOption {
	return func(option *searchOptionSet) {
		option.MaxDistance = maxDistance
	}
}

// WithLimit sets the maximum number of results that will be returned.
// A 0 value means no limit.
func WithLimit(limit int) SearchOption {
	return func(option *searchOptionSet) {
		option.Limit = limit
	}
}
