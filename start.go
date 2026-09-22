/*
format:
	title
	constraints
	question - starter 
	question - conditions/constraints
	(questions may include images)
	abbrivations - * explaination
	input - constraint 
	output - constraint
	input 
	output 
	testcase explaination (might inclued images)
*/
package main

import(
	"fmt"
	//"os"
	"time"
	"strings"
	"github.com/gocolly/colly/v2"
)
type problem struct {
	requeste_at time.Time //time of request for the record V^_^V
	title string
	time string
	space string
	problem_header string //handle for images
  problem_body string //handle here too 
	problem_footer string //and here
	input string 
	output string
	sample_tests string
}
func main(){
	var fetched_problem problem
	c:= colly.NewCollector(
		colly.AllowedDomains("codeforces.com"),
		colly.Async(true),
	)
	c.OnHTML(".problem-statement", func(e *colly.HTMLElement){
	fetched_problem.title = e.ChildText(".header .title")
  // fetched_problem.time = e.ChildText("*> .header .time-limit") 
  // fetched_problem.space = e.ChildText("* >.header .memory-limit")
	fmt.Println(fetched_problem.time)
	c.OnHTML(".header .time-limit", func(el *colly.HTMLElement){
		// time-limit has property class that has "property-title" that "time limit per test"
		// just ugly formatting trying to format without the grand child element 
		clone := el.DOM.Clone()
		clone.Children().Remove()
		fetched_problem.time = strings.TrimSpace(clone.Text())
	})
	})
	c.Visit("https://codeforces.com/problemset/problem/2264/D")
  c.Wait()
}

