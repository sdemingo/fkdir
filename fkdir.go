package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var words = []string{
	"ad",
	"adipisicing",
	"aliqua",
	"aliquip",
	"amet",
	"anim",
	"aute",
	"cillum",
	"commodo",
	"consectetur",
	"consequat",
	"culpa",
	"cupidatat",
	"deserunt",
	"do",
	"dolor",
	"dolore",
	"duis",
	"ea",
	"eiusmod",
	"elit",
	"enim",
	"esse",
	"est",
	"et",
	"eu",
	"ex",
	"excepteur",
	"exercitation",
	"fugiat",
	"id",
	"in",
	"incididunt",
	"ipsum",
	"irure",
	"labore",
	"laboris",
	"laborum",
	"Lorem",
	"magna",
	"minim",
	"mollit",
	"nisi",
	"non",
	"nostrud",
	"nulla",
	"occaecat",
	"officia",
	"pariatur",
	"proident",
	"qui",
	"quis",
	"reprehenderit",
	"sint",
	"sit",
	"sunt",
	"tempor",
	"ullamco",
	"ut",
	"velit",
	"veniam",
	"voluptate",
}

func RandomParagraph(min, max int) string {
	text := ""
	if min >= max {
		min = 0
	}
	nwords := rand.Intn(max-min) + min
	for i := 0; i < nwords; i++ {
		w := words[rand.Intn(len(words)-1)]
		if i == 0 {
			w = strings.Title(w)
		}
		text = text + " " + w
	}

	return strings.Trim(text, " ") + "."
}

func RandomDate() time.Time {
	min := time.Date(2016, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	text := RandomParagraph(50, 100)
	fmt.Println(text)
}
