package main

var (
	// ThesaurusURL returns json
	ThesaurusURL string = "https://words.bighugelabs.com/api/2/b6345742666ef3aef39d51c44710b272/%s/json"
	oxfordAPI    string = "https://od-api.oxforddictionaries.com/api/v2/words/en-us/%s"
	// DictionaryAPI response definition of word
	DictionaryAPI string = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=e1513aa1-d1c4-46a7-ad82-c6bd7c2eddd2"
)

// Definition struct of json response
type Definition []struct {
	Meta struct {
		ID        string   `json:"id"`
		UUID      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound,omitempty"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Ins []struct {
		Il  string `json:"il"`
		Ifc string `json:"ifc"`
		If  string `json:"if"`
	} `json:"ins"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Date     string   `json:"date"`
	Shortdef []string `json:"shortdef"`
}

// func getWord(word string) []byte {
// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", fmt.Sprintf(oxfordAPI, word), nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	req.Header.Set("app_id", "32e73c16")
// 	req.Header.Set("app_key", "a9a3122a4f8f582f9d091f8dde2d8f5f")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(req.Header)

// 	return body
// }
