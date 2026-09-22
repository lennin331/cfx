package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

// convertCodeforcesLaTeX converts $$$ delimiter to standard LaTeX $ delimiters
func convertCodeforcesLaTeX(text string) string {
	// Replaces all Codeforces $$$ tags with standard LaTeX $ tags
	return strings.ReplaceAll(text, "$$$", "$")
}

func main() {
	targetURL := "https://codeforces.com/problemset/problem/2264/A"

	if len(os.Args) > 1 {
		targetURL = os.Args[1]
	}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"),
	)

	c.OnHTML(".problem-statement", func(e *colly.HTMLElement) {
		// Title and Metadata
		title := convertCodeforcesLaTeX(strings.TrimSpace(e.ChildText(".header .title")))
		timeLimit := convertCodeforcesLaTeX(strings.TrimSpace(e.ChildText(".header .time-limit")))
		memoryLimit := convertCodeforcesLaTeX(strings.TrimSpace(e.ChildText(".header .memory-limit")))

		fmt.Printf("==================================================\n")
		fmt.Printf("TITLE: %s\n", title)
		fmt.Printf("%s\n", timeLimit)
		fmt.Printf("%s\n", memoryLimit)
		fmt.Printf("==================================================\n\n")

		// Problem Description
		fmt.Println("--- PROBLEM STATEMENT ---")
		e.ForEach(".header + div p", func(_ int, el *colly.HTMLElement) {
			text := convertCodeforcesLaTeX(strings.TrimSpace(el.Text))
			if text != "" {
				fmt.Println(text)
				fmt.Println()
			}
		})

		// Input Specification
		inputSpec := convertCodeforcesLaTeX(strings.TrimSpace(e.ChildText(".input-specification")))
		if inputSpec != "" {
			fmt.Println("--- INPUT SPECIFICATION ---")
			fmt.Println(strings.TrimPrefix(inputSpec, "Input"))
			fmt.Println()
		}

		// Output Specification
		outputSpec := convertCodeforcesLaTeX(strings.TrimSpace(e.ChildText(".output-specification")))
		if outputSpec != "" {
			fmt.Println("--- OUTPUT SPECIFICATION ---")
			fmt.Println(strings.TrimPrefix(outputSpec, "Output"))
			fmt.Println()
		}

		// Sample Tests
		fmt.Println("--- SAMPLE TESTS ---")
		e.ForEach(".sample-test .sample-test", func(i int, sample *colly.HTMLElement) {
			input := strings.TrimSpace(sample.ChildText(".input pre"))
			output := strings.TrimSpace(sample.ChildText(".output pre"))

			fmt.Printf("Example %d:\n", i+1)
			fmt.Printf("Input:\n%s\n", input)
			fmt.Printf("Output:\n%s\n\n", output)
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error fetching page:", err)
	})

	c.Visit(targetURL)
}
