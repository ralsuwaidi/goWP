package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-wordwrap"
	termbox "github.com/nsf/termbox-go"
)

var redditURL string = "https://www.reddit.com/r/WritingPrompts/.json?limit=20&"
var topWeekURL string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=week"
var topMonthURL string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=month"
var topYearURL string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=year"
var terminalWidth int

// getResponse returns http GET request in bytes
func getResponse(url, userAgent string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func getPosts(byt []byte) Posts {
	res := Posts{}

	json.Unmarshal(byt, &res)

	return res
}

func getComments(byt []byte) Comments {
	res := Comments{}

	if err := json.Unmarshal(byt, &res); err != nil {
		panic(err)
	}
	return res
}

type writingPrompt struct {
	title string
	date  time.Time
	story string
	award bool
}

func makePrompt(posts Posts, number int) writingPrompt {
	stickied := 0
	for posts.Data.Children[0+stickied].Data.Stickied {
		stickied++
	}

	// skip stickied
	number = number + stickied
	// get url
	url := posts.Data.Children[number].Data.URL + ".json"
	// comments from url
	commentsByt := getResponse(url, "Golang_Spider_Bot/3.05")
	comments := getComments(commentsByt)[1].Data.Children
	story := comments[1].Data.Body
	wp := writingPrompt{
		title: posts.Data.Children[number].Data.Title,
		date:  time.Unix(int64(posts.Data.Children[number].Data.CreatedUtc), 0),
		story: story,
	}

	if posts.Data.Children[number].Data.AllAwardings != nil {
		if len(posts.Data.Children[number].Data.AllAwardings) > 0 {
			wp.award = true
		}
	}

	return wp
}

func savePrompt(wp writingPrompt) {
	wpDump := "\n" + "\n" + wp.title + "\n" + wp.story + "\n"
	f, err := os.OpenFile("saved_wp.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(wpDump); err != nil {
		panic(err)
	}
}

// return sar if awarded
func award(wp writingPrompt) string {
	if wp.award {
		return "[*]"
	}
	return ""
}

// returns posts depending on sorting order
func sortWP(sort string) Posts {
	var posts Posts

	switch sort {
	case "week":
		response := getResponse(topWeekURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top week]")
	case "month":
		response := getResponse(topMonthURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top month]")
	case "year":
		response := getResponse(topYearURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to top year]")
	default:
		response := getResponse(redditURL, "Golang_Spider_Bot/3.0")
		posts = getPosts(response)
		fmt.Println("[changed to hot]")
	}
	return posts
}

func printWrapped(text string) {
	if terminalWidth == 0 {
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		w, _ := termbox.Size()
		termbox.Close()
		terminalWidth = w
	}

	wrapped := wordwrap.WrapString(text, uint(terminalWidth))
	fmt.Println(wrapped)
}

func main() {
	promptInt := new(int)

	// get wp
	response := getResponse(redditURL, "Golang_Spider_Bot/3.0")
	posts := getPosts(response)
	wp := makePrompt(posts, *promptInt)

	// print title and get user input
	reader := bufio.NewReader(os.Stdin)
	printWrapped("\n" + award(wp) + wp.title + "\n")
	fmt.Print("> Read? [y/N]: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// loop titles until 'y' selected
	for strings.TrimSpace(userInput) != "y" {
		*promptInt++

		// sort time if input
		if strings.TrimSpace(userInput) == "week" {
			posts = sortWP("week")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "month" {
			posts = sortWP("month")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "year" {
			posts = sortWP("year")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "hot" {
			posts = sortWP("week")
			*promptInt = 0
		}

		wp = makePrompt(posts, *promptInt)
		printWrapped("\n" + award(wp) + wp.title + "\n")
		fmt.Print("> Read? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

	}

	// loop over story text
	splitStory := strings.Split(wp.story, "\n\n")
	splitStoryPt := &splitStory

	fmt.Println("\n ")
	saved := new(bool)
	for i := 0; i < len(*splitStoryPt); i++ {
		printWrapped((*splitStoryPt)[i])
		reader = bufio.NewReader(os.Stdin)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(userInput) == "s" && *saved == false {
			savePrompt(wp)
			fmt.Println("[saved to 'saved_wp.txt']\n ")
			*saved = true
		} else if strings.TrimSpace(userInput) == "s" && *saved == true {
			fmt.Println("[already saved']\n ")
		}
	}

	fmt.Print("> Done! Want to save? [y/N]: ")
	userInput, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if strings.TrimSpace(userInput) == "y" {
		savePrompt(wp)
		fmt.Println("[saved to 'saved_wp.txt']\n ")
	}

}
