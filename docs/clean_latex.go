//parsing the fetched text to be adapted in the terminal env
import "strings"

func cleanCodeforcesText(text string) string {
	// Remove MathJax markers
	text = strings.ReplaceAll(text, "$$$", "")
	
	// Convert common LaTeX symbols to Unicode for terminal readability
	replacements := map[string]string{
		"\\le": "≤",
		"\\ge": "≥",
		"\\neq": "≠",
		"\\times": "×",
		"\\dots": "...",
		"\\in": "in",
	}
	
	for latex, unicode := range replacements {
		text = strings.ReplaceAll(text, latex, unicode)
	}
	
	return text
}
