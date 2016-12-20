package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
)

var dbURL = "http://172.17.0.4:8099"

func main() {
	
    r := mux.NewRouter()
    r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Welcome to Mainsail Message Lookup Tool! ")
    })
	r.HandleFunc("/querybyno/{msgNo}", displayMsg)
	r.HandleFunc("/queryall", displayMsg)
	http.ListenAndServe(":8100", r)
}


func displayMsg(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(dbURL + r.URL.Path  ) 
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w,string(body))
}
