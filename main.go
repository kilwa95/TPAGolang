package main


import (
	"net/http"
	"time"
	"fmt"
	"os"
    "io/ioutil"

)


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
	    f, err := os.OpenFile("./save.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = fmt.Fprintf(f, "%s: %s\n", req.PostForm["author"][0], req.PostForm["entry"][0])

		if err != nil {
		fmt.Println(err)
		f.Close()
		return
		}
		err = f.Close()
		if err != nil {
		fmt.Println(err)
		return
		}
		fmt.Fprintf(w, "file appended successfully")
	}
}


func getListAuthHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		data, err := ioutil.ReadFile("save.data")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Contents of file:", string(data))
		fmt.Fprintf(w, string(data))

	}
}



func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/hello", createAuthHandler)
	http.HandleFunc("/entries", getListAuthHandler)
	http.ListenAndServe(":4567", nil)
}