package main


import (
	"net/http"
	"time"
	"fmt"
	"os"
	"bufio"

)

// type Auth struct {
//     name string 
//     text string 
// }



func timeHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	d := time.Date(2000, 9, 13, 11, 2, 0, 0, time.UTC)
    day := d.Day()
    minute := d.Minute()
	data := fmt.Sprintf("%vh%v", day, minute)
	fmt.Fprintf(w,data)
	}
}


func createAuthHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

	if err := req.ParseForm(); err != nil {
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
		return
		}
		saveFile, err := os.OpenFile("./save.data", os.O_RDWR|os.O_CREATE, 0644)
		defer saveFile.Close()
		w := bufio.NewWriter(saveFile)

		if err == nil {
			fmt.Fprintf(w, "%s: %s\n", req.PostForm["author"], req.PostForm["entry"])
		}
		w.Flush()
	}
}


func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/hello", createAuthHandler)
	http.ListenAndServe(":4567", nil)
}