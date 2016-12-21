package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os"
)

func main() {
	
    r := mux.NewRouter()
    r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Welcome to Mainsail Message Lookup Tool! " + os.Getenv("DBURL"))
    })
	r.HandleFunc("/querybyno/{msgNo}", displayMsg)
	r.HandleFunc("/queryall", displayMsg)
	http.ListenAndServe(":8100", r)
}


func displayMsg(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(os.Getenv("DBURL") + r.URL.Path) 
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w,string(body))
}
