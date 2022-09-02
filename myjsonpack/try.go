package myjsonpack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Myjson struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Unmars Myjson

func JsonReqHandler() (Myjson, error) {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("ERROR!...")
		log.Fatal(err)
	}
	log.Print(response.StatusCode)

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Myjson{}, err
	}

	byt := []byte(bytes)
	err = json.Unmarshal(byt, &Unmars)
	if err != nil {
		log.Fatal(err)
	}

	return Unmars, nil
}

func (s Myjson) PostJson() {
	s.Completed = false
	s.ID = 1
	s.Title = "perfect"
	s.UserID = 1
	send, err := json.Marshal(s)
	if err != nil {
		fmt.Println("an error occured")
		return
	}
	gonder := bytes.NewBuffer(send)
	resp, _ := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json", gonder)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}

func My() {
	toPost := Myjson{}
	toPost.PostJson()

}
