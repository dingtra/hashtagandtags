package hashtag


import (
	"strings"
)

func EmbedTagsInBio ( text string ) string {
	tag := strings.Split(text, " ")
	track := map[int]string{}

	if len(text) > 0 {
		for i, k := range tag {

			if len(k) > 1 {
				if k[0:1] == "@" {
					track[i] = "<a href='/"+k[1:]+"'>"+k+"</a>"
				}
			}
		}
	}

	if len(track) > 0 {
		for i, k := range track {

			tag[i] = k
		}
	}

	return strings.Join(tag, " ")
}