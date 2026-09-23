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
	//"strings"
	"regexp"
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
	fetched_problem.title = e.ChildText(".header .title") // title - variable 

  fetched_problem.space = e.ChildText(".header .memory-limit .property-title") 
  fetched_problem.time = e.ChildText(".header .time-limit .property-title") 

	memory_value:= e.ChildText(".header .memory-limit ")
	time_value := e.ChildText(".header .time-limit ")
	//regex to get only the number for the time_value - varuable
	reg := regexp.MustCompile(`[^\d.]`)
	time_value = reg.ReplaceAllString(time_value, "")
	memory_value = reg.ReplaceAllString(memory_value, "")






	//printables are kept seperately 
	fmt.Println(fetched_problem.title)
	fmt.Println(fetched_problem.time + ": " + time_value + "s")
	fmt.Println(fetched_problem.space + ": " + memory_value + "Mb")
	})
	c.Visit("https://codeforces.com/problemset/problem/2264/D")
  c.Wait()
}

