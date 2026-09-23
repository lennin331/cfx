package main
import(
	"fmt"
	//"os"
	"time"
	"strings"
	"regexp"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)
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
type problem struct {
	requeste_at time.Time //time of request for the record V^_^V
	title string
	time string
	space string
	problem_text string //handle for images
  problem_input_specification string //handle here too 
	problem_output_specification string //and here
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

	
e.DOM.Find(".header ~ div").Each(func(_ int, s * goquery.Selection){
		if s.HasClass("input-specification") || s.HasClass("output-specification") || s.HasClass("sample-tests"){
			return
		}
		s.Find("p").Each(func(_ int, p* goquery.Selection){
			text:= strings.TrimSpace(p.Text())
			if text != "" {
				fetched_problem.problem_text += text +"\n"
			}
		})
		fetched_problem.problem_text += "\n\n" // this apparently handles the note part at the bottom of the page
	})

	//printables are kept seperately 
	fmt.Println(fetched_problem.title)
	fmt.Println(fetched_problem.time + ": " + time_value + "s")
	fmt.Println(fetched_problem.space + ": " + memory_value + "Mb")
	fmt.Println(fetched_problem.problem_text)
	})
	c.Visit("https://codeforces.com/problemset/problem/2264/D")
  c.Wait()
}

