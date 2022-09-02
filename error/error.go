package error

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Myhttp() string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	fmt.Println()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bd := string(body)
	return bd
}
