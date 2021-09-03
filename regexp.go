package main

import (
	"fmt"
	"regexp"
)



func main() {
	
	var regex *regexp.Regexp=regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("abdul")) // true
	fmt.Println(regex.MatchString("acdul")) // true
	fmt.Println(regex.MatchString("aDdul")) // false

	search:=regex.FindAllString("abdul abdul abdul abdul",2)
	fmt.Println(search)
}