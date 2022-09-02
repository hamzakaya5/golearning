package main

import (
	"fmt"
	"net/http"
)

func main() {

	/*---------FILE HANDLING---------*/

	// file, err := os.Create("file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	file.WriteString("naber mudur")
	// 	file.Close()
	// }
	// data, err := ioutil.ReadFile("file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	in := string(data)
	// 	fmt.Println(in)
	// }

	/*----------HTTP SERVER----------*/

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello\n")
	fmt.Println("w = ", w, "\nreq = ", req)

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Println("w = ", w, "\nreq = ", req)
		}
	}
}
