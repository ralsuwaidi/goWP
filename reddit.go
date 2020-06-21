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
)

var (
	redditURL     string = "https://www.reddit.com/r/WritingPrompts/.json?limit=20&"
	topWeekURL    string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=week"
	topMonthURL   string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=month"
	topYearURL    string = "https://www.reddit.com/r/WritingPrompts/top/.json?t=year"
	terminalWidth int
	promptInt     *int = new(int)
	posts         Posts
	userInput     string
	err           error
	wp            writingPrompt
)

// GetResponse returns http GET request in bytes
func GetResponse(url, userAgent string) []byte {
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
	var commentsByt []byte
	// find stickied
	stickied := 0
	for posts.Data.Children[0+stickied].Data.Stickied {
		stickied++
	}

	// skip stickied
	number = number + stickied
	// get url
	url := posts.Data.Children[number].Data.URL + ".json"
	// comments from url
	commentsByt = GetResponse(url, "Golang_Spider_Bot/3.05")

	// check if comments exist
	// skip if there is no comment
	for len(getComments(commentsByt)[1].Data.Children) < 2 {
		number++
		url := posts.Data.Children[number].Data.URL + ".json"
		commentsByt = GetResponse(url, "Golang_Spider_Bot/3.05")
		*promptInt++
	}

	// get comments and make a wp struct
	comments := getComments(commentsByt)[1].Data.Children
	story := comments[1].Data.Body
	wp := writingPrompt{
		title: posts.Data.Children[number].Data.Title,
		date:  time.Unix(int64(posts.Data.Children[number].Data.CreatedUtc), 0),
		story: story,
	}

	// add award
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

// return star if awarded
func award(wp writingPrompt) string {
	if wp.award {
		return "[*]"
	}
	return ""
}

func main() {

	var definition Definition
	// get posts
	response := GetResponse(redditURL, "Golang_Spider_Bot/3.0")
	posts = getPosts(response)

	// print title and get user input
	reader := bufio.NewReader(os.Stdin)

	// loop titles until story is selected
	for strings.TrimSpace(userInput) != "y" {
		wp = makePrompt(posts, *promptInt)
		PrintWrapped("\n" + award(wp) + wp.title + "\n")
		fmt.Print("> Read? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(userInput) == "exit" {
			os.Exit(0)
		}

		*promptInt++

		// sort time if input
		if strings.TrimSpace(userInput) == "week" {
			posts = SortWP("week")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "month" {
			posts = SortWP("month")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "year" {
			posts = SortWP("year")
			*promptInt = 0
		} else if strings.TrimSpace(userInput) == "hot" {
			posts = SortWP("hot")
			*promptInt = 0
		}

	}

	// loop over story text
	splitStory := strings.Split(wp.story, "\n\n")
	splitStoryPt := &splitStory

	fmt.Println("\n ")
	saved := new(bool)
	for i := 0; i < len(*splitStoryPt); i++ {
		PrintWrapped((*splitStoryPt)[i])
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
		} else if strings.Contains(strings.TrimSpace(userInput), "def") {
			// split
			word := strings.Split(strings.TrimSpace(userInput), " ")
			body := GetResponse(fmt.Sprintf(DictionaryAPI, word[1]), "This is a test")
			json.Unmarshal(body, &definition)
			fmt.Println()
			fmt.Println(definition[0].Shortdef)
			userInput, _ = reader.ReadString('\n')
		} else if strings.TrimSpace(userInput) == "exit" {
			os.Exit(0)
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
