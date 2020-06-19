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

	return writingPrompt{
		title: posts.Data.Children[number].Data.Title,
		date:  time.Unix(int64(posts.Data.Children[number].Data.CreatedUtc), 0),
		story: story,
	}
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

func main() {
	promptInt := new(int)

	// get wp
	response := getResponse("https://www.reddit.com/r/WritingPrompts/.json?limit=20", "Golang_Spider_Bot/3.0")
	posts := getPosts(response)
	wp := makePrompt(posts, *promptInt)
	wpPt := &wp

	// get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + wp.title + "\n")
	fmt.Print("Read? [y/N]: ")
	char, _, _ := reader.ReadRune()

	// loop titles until 'y' selected
	for char != 'y' {
		*promptInt++
		*wpPt = makePrompt(posts, *promptInt)
		fmt.Println("\n" + wp.title + "\n")
		fmt.Print("Read? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		char, _, _ = reader.ReadRune()
	}

	// loop over story text
	splitStory := strings.Split(wp.story, "\n\n")
	fmt.Println("\n ")
	saved := new(bool)
	for i := 0; i < len(splitStory); i++ {
		fmt.Println(splitStory[i])
		reader = bufio.NewReader(os.Stdin)
		char, _, _ = reader.ReadRune()
		if char == 's' && *saved == false {
			savePrompt(wp)
			fmt.Println("[saved to 'saved_wp.txt']\n ")
			*saved = true
		} else if char == 's' && *saved == true {
			fmt.Println("[already saved']\n ")

		}
	}
}
