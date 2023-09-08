package packagego

import (
	"fmt"
	"regexp"
)

func MainPackageRegexp() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z]*)o")

	fmt.Println(regex.MatchString("ejo"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("etso"))

	fmt.Println(regex.FindAllString("eko eka elo eto eki", 2))
}
