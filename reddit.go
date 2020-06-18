package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

	number = number + stickied
	url := posts.Data.Children[number].Data.URL + ".json"
	commentsByt := getResponse(url, "Golang_Spider_Bot/3.0")
	comments := getComments(commentsByt)[1].Data.Children
	story := comments[1].Data.Body

	return writingPrompt{
		title: posts.Data.Children[number].Data.Title,
		date:  time.Unix(int64(posts.Data.Children[number].Data.CreatedUtc), 0),
		story: story,
	}
}

func main() {

	response := getResponse("https://www.reddit.com/r/WritingPrompts/.json?limit=20", "Golang_Spider_Bot/3.0")
	posts := getPosts(response)

	var userInt int
	fmt.Print("WP number: ")
	fmt.Scan(&userInt)

	wp := makePrompt(posts, userInt)
	fmt.Println(wp.title)
	fmt.Println(wp.story)
}
