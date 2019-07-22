package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)
// global variable

var data wStruct
{
	ImageURL string `json:"imageUrl"`
}

// made a simple struct that uses the ImageUrl received in json Payload. You can add these 
// to the struct to get more information.
//type test struct {
//	EvalMatches []struct {
//		Value  int         `json:"value"`
//		Metric string      `json:"metric"`
//		Tags   interface{} `json:"tags"`
//	} `json:"evalMatches"`
//	ImageURL string `json:"imageUrl"`
//	Message  string `json:"message"`
//	RuleID   int    `json:"ruleId"`
//	RuleName string `json:"ruleName"`
//	RuleURL  string `json:"ruleUrl"`
//	State    string `json:"state"`
//	Title    string `json:"title"`
//}



// main that starts 8045/, and 8045 /alert

func main() {
	log.Println("server started")
	http.HandleFunc("/", handleWebhook)
	http.HandleFunc("/alert", handleMainPage)
	log.Fatal(http.ListenAndServe(":8045", nil))
}

// takes json payload and unmarshals the data into viable information

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return
	}
	fmt.Println("got webhook payload: ", data)

}

// creates a main page and should take our variable and pass it to the "homepage.html"

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	fmt.Println(data.ImageURL)
	err = tmpl.Execute(w, data) 
	if err != nil {
		log.Print("template executing error: ", err)
	}

}
