package shared

import (
	"strings"
)

// Normalize trims and lowers a string
func Normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// ContainsDisqualifier checks if a search term contains any disqualifying word/phrase
func ContainsDisqualifier(term string, disqualifiers []string) bool {
	term = Normalize(term)
	for _, dq := range disqualifiers {
		if strings.Contains(term, dq) {
			return true
		}
	}
	return false
}

// Deduplicate returns only unique strings from the given list
func Deduplicate(terms []string) []string {
	seen := make(map[string]bool)
	unique := []string{}
	for _, t := range terms {
		lowered := strings.ToLower(t)
		if !seen[lowered] {
			seen[lowered] = true
			unique = append(unique, t)
		}
	}
	return unique
}
