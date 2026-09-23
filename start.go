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
	request_at time.Time //time of request for the record V^_^V
	title string
	time string
	space string
	problem_text string //handle for images
	input []string 
	output []string
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
		if s.HasClass("input-specification") || s.HasClass("output-specification") || s.HasClass("sample-tests") ||(s.HasClass("note")){
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

	fetched_problem.request_at = time.Now()
	})
c.OnHTML(".sample-tests .sample-test", func(e *colly.HTMLElement) {
		
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
			fetched_problem.input = append(fetched_problem.input, singleInput)
		}
		if singleOutput != "" {
			fetched_problem.output = append(fetched_problem.output, singleOutput)
		}
		
	})



	start_time := time.Now()
	c.Visit("https://codeforces.com/problemset/problem/2264/D")
	c.Wait()
	fmt.Println("Test Cases : " + fetched_problem.input[0])
	//passing the first element resolves it as a string when the whole structure is []string, can't really add the "Input :" after the first element, but is good enough to format with testcase follwed by the input values
	for i:= 0; i<len(fetched_problem.input); i++ {

		// fmt.Print(fetched_problem.input[i])
		// How the fuck does this even work, somehow this is formatting perfectly 
		// fml
		fmt.Print("\n\nOutput :\n")
		fmt.Println(fetched_problem.output[i])
	 // fmt.Println(fetched_problem.output[i])
	}
	 
	fmt.Println(fmt.Sprintf("\nFetched : [ %.6s ms ]\n", fetched_problem.request_at.Sub(start_time)))
}
