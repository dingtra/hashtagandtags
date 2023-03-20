package hashtag


import (
	"strings"
	"regexp"
)


func TagSomeOne (writeup string ) {


	if len(writeup) > 0 {
		for _, k := range strings.Split(writeup, " ") {

			if k[0:1] == "@" {
				checkname := regexp.MustCompile(`[^A-Za-z0-9]`)
				resultname := checkname.ReplaceAllString(k[1:], "")

			}
		}
	}
}