package hashtag


import (
	"strings"
	"html/template"
	// "fmt"
)


func SeeMore(item, id string) string {

	bks :=  strings.Split(item, " ")
	before := []string{}
	after := []string{}
	counter := 0

	if len(bks) > 0 {

		for _, item :=  range bks {

			if len(before) < 20 {
				before = append(before, item)
			}
			counter +=1

			if len(before) == 20 {

				if counter > 20 {
					after = append(after, item)
				}
			}
		}
	}

	if len(before) > 0 && len(after) > 0 {
		// do the math here

		return `
		    <div style="flex-direction:column;" class="xklop090">
			    <div class="xklop090-chld">
			        <div class='xklop091`+id+`'> `+strings.Join(before, " ")+` </div>
				</div>
				<div style="opacity:0.7;font-size:15px;" all-data="`+template.HTMLEscapeString(strings.Join(before, " "))+" "+template.HTMLEscapeString(strings.Join(after, " "))+`" all-less='`+template.HTMLEscapeString(strings.Join(before, " "))+`' id='xklop092m00`+id+`' target="more"><div style="pointer-events:none;">... See more <i style="height:15px;width:15px;position:relative;top:2px;" data-feather='chevron-up'></i> </div></div>
			</div>
		`
	}else{

		if len(before) > 0 {
			return `
		    <div class="xklop090">
			    <div class="xklop091"> `+strings.Join(before, " ")+` </div>
			</div>
			`
		}
	}

	return ""
}


func SeeMorex (text, id string ) string {

	var writeup string

	if len(text) > 0 {

		counter := []string{}
		var checkcounter bool

		if len(text) > 20 {

			brks := strings.Split(text, " ")

			for _, k := range brks {

				if len(k) > 1 {

					counter = append(counter, k)
					
					if len(counter)   == 15 {
						checkcounter = true
						counter = append(counter, "<div style='display:none;' id='infoxr"+id+"' class='writeup-class'><span class='see-more-span'> "+k)
					}
				}
			}

			if checkcounter == true {
				counter = append(counter, "</span></div> <div id='clse"+id+"'><span class='see-more-tag-class' id='infoxr-more' kl='"+id+"' style='cursor: pointer;color:blue;opacity:0.8;'>See more  <i data-feather='chevron-up'></i></span></div>")
			}
			writeup = strings.Join(counter, " ")
		}else{
			writeup = text
		}

	}else{
		writeup = text
	}

	return writeup

	
}