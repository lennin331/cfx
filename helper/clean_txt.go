package helper 
import(
	"strings"
)
func CleanProblemText(text string) string {
	// Remove MathJax markers
	text = strings.ReplaceAll(text, "$$$", "")
	
	// Convert common LaTeX symbols to Unicode for terminal readability
	replacements := map[string]string{
		"\\le": "≤",
		"\\ge": "≥",
		"\\neq": "≠",
		"\\times": "×",
		"\\dots": "...",
		"\\ldots": "...",
		"\\in": "in",
		"\\lt" : "<",
		"\\gcd" : "gcd",
		//s_i=matht{some_text}
		//\texttrm{some_text}
	}
	
	for latex, unicode := range replacements {
		text = strings.ReplaceAll(text, latex, unicode)
	}
	
	return text
}
