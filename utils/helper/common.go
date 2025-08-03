package helper

import "strings"

func CleanString(input string) string {
	// Define replacements in a map for better maintainability
	replacements := map[string]string{
		"\n":     "",  // newline
		"\t":     "",  // tab
		"\u0026": "and", // ampersand
		"\u2011": "-", // non-breaking hyphen
		"\u00a0": " ", // non-breaking space
		// Add more replacements as needed
	}

	// Apply all replacements
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	// Trim leading/trailing whitespace and return
	return strings.TrimSpace(input)
}
