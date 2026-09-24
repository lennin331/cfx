//Input from the console on  
//contest and problem ID
package helper

import (
	"fmt"
	"regexp"
)
func Get_id() string{
	var IsAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	var IsDigit = regexp.MustCompile(`^[0-9]+$`).MatchString
  var target_contest string 
	var target_problem string
	_, err := fmt.Scan(&target_contest, &target_problem)
	if err != nil || !IsAlpha(target_problem) || !IsDigit(target_contest)  {
		fmt.Println("Couldn't Find the Problem\n <contest-number> <problem-id>\n %s", err)
		return ""
	}
	
	return target_contest + "/" + target_problem
}
