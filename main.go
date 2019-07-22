package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var data wStruct

// made a simple struct that uses the ImageUrl received in json Payload
{
	ImageURL string `json:"imageUrl"`
}

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
