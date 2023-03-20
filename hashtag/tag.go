package hashtag


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"regexp"
	"fmt"
)


func TagSomeOne (shared string, items primitive.M, users []primitive.M ) string {

	postid := ""

	fmt.Println(items["body"].(string), " BODY")

	if len(shared) > 0 {
		// post is shared used share id for => show more and => show less
		postid = shared
	}else{
		// post not shared used postid for => show more and => show less
		postid = items["_id"].(primitive.ObjectID).Hex()
	}

	grabtags := make(map[string]int)
	result := make(map[string]string)
	hash := make(map[int]string)
	link := make(map[int]string)

	if items["body"] != nil && len(items["body"].(string)) > 0 {
		for i, k := range strings.Split(SeeMore(items["body"].(string), postid), " ") {

			if len(k) > 1 {
				if k[0:1] == "@" {
					fmt.Println(k, " JJJJ")
					checkname := regexp.MustCompile(`[^A-Za-z0-9]`)
					resultname := checkname.ReplaceAllString(k[1:], "")
	
					grabtags[resultname] = i
					result[resultname] = resultname
	
				}else if k[0:1] == "#" {

					hash[i] = k

				}else if len(k) > 3 {
					if k[0:4] == "http" {
						link[i] = k
					}

					if k[0:3] == "www" {
						link[i] = k
					}

				}
			}
		}

	}

	finalwrap := make(map[int]string)


	for _, item := range users {

		if len(result[item["username"].(string)]) > 0 {

			finalwrap[ grabtags[item["username"].(string)] ] = item["username"].(string)
		}
	}

	tagresult := strings.Split(SeeMore(items["body"].(string), postid), " ")


	if len(hash) > 0 || len(tagresult) > 0 || len(link) > 0 {
		for i, item := range finalwrap {
	
			tagresult[i] = "<a href='/"+item+"'>"+item+"</a>"
		}
	
		for i, item := range hash {
	
			tagresult[i] = "<a style='text-decoration:none;' href='/hashtag/preview/?search="+item[1:]+"'>"+item+"</a>"
		}

		for i, item := range link {
			tagresult[i] = "<a href='"+item+"'>"+item+"</a>"
		}

	}


	return strings.Join(tagresult, " ")
}