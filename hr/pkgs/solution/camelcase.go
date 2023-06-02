package solution

import (
	"log"
	"regexp"
)

func Camelcase(s string) int32 {
	regex := "[A-Z]"

	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	words := re.Split(s, -1)
	return int32(len(words))
}