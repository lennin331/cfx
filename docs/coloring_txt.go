//Ref - https://zetcode.com/golang/terminal-colour/

package main

import (
    "fmt"
		"github.com/logrusorgru/aurora/v4"
)

func main() {

    msg := "an old falcon"

    reset := "\033[0m"
    bold := "\033[1m"
    underline := "\033[4m"
    strike := "\033[9m"
    italic := "\033[3m"

    cRed := "\033[31m"
    cGreen := "\033[32m"
    cYellow := "\033[33m"
    cBlue := "\033[34m"
    cPurple := "\033[35m"
    cCyan := "\033[36m"
    cWhite := "\033[37m"

    fmt.Println(msg)

    fmt.Println(cRed + msg)
    fmt.Println(cGreen + msg)
    fmt.Println(cYellow + msg)
    fmt.Println(cBlue + msg)
    fmt.Println(cPurple + msg)
    fmt.Println(cWhite + msg)
    fmt.Println(cCyan + msg + reset)

    fmt.Println(bold + msg)
    fmt.Println(italic + msg + reset)
    fmt.Println(strike + msg + reset)
    fmt.Println(underline + msg + reset)
    fmt.Println(msg)

    msg = " another old falcon "

    fmt.Println(aurora.Gray(0, msg).BgGray(23))
    fmt.Println(aurora.Gray(3, msg).BgGray(19))
    fmt.Println(aurora.Gray(7, msg).BgGray(15))
    fmt.Println(aurora.Gray(11, msg).BgGray(13))
    fmt.Println(aurora.Gray(15, msg).BgGray(7))
    fmt.Println(aurora.Gray(19, msg).BgGray(3))
    fmt.Println(aurora.Gray(23, msg).BgGray(0))

		
}
