package main

import (
	"fmt"
	"regexp"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	termbox "github.com/nsf/termbox-go"
)

var (
	// ThesaurusURL returns json
	ThesaurusURL string = "https://words.bighugelabs.com/api/2/b6345742666ef3aef39d51c44710b272/%s/json"
	oxfordAPI    string = "https://od-api.oxforddictionaries.com/api/v2/words/en-us/%s"
	// DictionaryAPI response definition of word
	DictionaryAPI string = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=e1513aa1-d1c4-46a7-ad82-c6bd7c2eddd2"
)

// PrintWrapped prints wrapped text
func PrintWrapped(text string) {
	if terminalWidth == 0 {
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		w, _ := termbox.Size()
		termbox.Close()
		terminalWidth = w
	}

	//wrapped := wordwrap.WrapString(text, uint(terminalWidth))
	result := markdown.Render(text, terminalWidth, 0)
	fmt.Println(string(result))
}

// SortWP changes sort order and refreshes the title list
func SortWP(sort string) Posts {
	var posts Posts

	switch sort {
	case "week":
		response := GetResponse(topWeekURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top week]")
	case "month":
		response := GetResponse(topMonthURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top month]")
	case "year":
		response := GetResponse(topYearURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top year]")
	default:
		response := GetResponse(redditURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to hot]")
	}
	return posts
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func findLinks(t string) map[string]string {
	m := make(map[string]string)
	var stringList []string

	MDRegex := `(?:__|[*#])|\[(.*?)\]\(.*?\)`
	rm, _ := regexp.Compile(MDRegex)

	match1, _ := regexp.MatchString(MDRegex, t)
	if match1 {
		MDList := rm.FindAllString(t, -1)
		for i := 0; i < len(MDList); i++ {
			if MDList[i] == "*" {
				continue
			}
			splitMD := strings.Split(MDList[i], "](")
			m[splitMD[0][1:]] = splitMD[1][:len(splitMD[1])-1]

		}
	}

	regexExpression := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&\/=]*)`
	r, _ := regexp.Compile(regexExpression)

	match, _ := regexp.MatchString(regexExpression, t)
	if match {
		stringList = r.FindAllString(t, -1)
	}

	for _, element := range m {
		if len(stringList) > 0 {
			for i := 0; i < len(stringList); i++ {
				if strings.Contains(element, stringList[i]) {

					stringList = remove(stringList, i)
				}
			}
		}
	}

	for i := 0; i < len(stringList); i++ {

		m[stringList[i]] = stringList[i]
	}

	return m
}
