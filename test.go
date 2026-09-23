package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	var inputs []string
	var outputs []string

	c.OnHTML(".problem-statement .sample-tests .sample-test", func(e *colly.HTMLElement) {
		// --- Extract Input ---
		inputPre := e.DOM.Find(".input pre")
		var singleInput string

		linesIn := inputPre.Find(".test-example-line")
		if linesIn.Length() > 0 {
			linesIn.Each(func(_ int, line *goquery.Selection) {
				singleInput += line.Text() + "\n"
			})
		} else {
			inputPre.Find("br").ReplaceWithHtml("\n")
			singleInput = inputPre.Text()
		}

		// --- Extract Output ---
		outputPre := e.DOM.Find(".output pre")
		var singleOutput string

		linesOut := outputPre.Find(".test-example-line")
		if linesOut.Length() > 0 {
			linesOut.Each(func(_ int, line *goquery.Selection) {
				singleOutput += line.Text() + "\n"
			})
		} else {
			outputPre.Find("br").ReplaceWithHtml("\n")
			singleOutput = outputPre.Text()
		}

		// Trim extra padding and store in slices
		singleInput = strings.TrimSpace(singleInput)
		singleOutput = strings.TrimSpace(singleOutput)

		if singleInput != "" {
			inputs = append(inputs, singleInput)
		}
		if singleOutput != "" {
			outputs = append(outputs, singleOutput)
		}
	})

	err := c.Visit("https://codeforces.com/problemset/problem/2264/C")
	if err != nil {
		log.Fatal(err)
	}

	// Print inputs and outputs paired one by one
	for i := 0; i < len(inputs); i++ {
		fmt.Printf("=== SAMPLE TEST #%d ===\n", i+1)
		
		fmt.Println("--- INPUT ---")
		fmt.Println(inputs[i])
		
		fmt.Println("\n--- OUTPUT ---")
		fmt.Println(outputs[i])
		
		fmt.Println("======================\n")
	}
}
