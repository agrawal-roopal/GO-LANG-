package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello World")
		d,err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "Ooops",http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))
			return

		}
		fmt.Fprintf(rw, "hello %s\n",d)
	})
	http.HandleFunc("/goodbye",func(http.ResponseWriter, *http.Request){
		log.Println("goodbye world")
	})
	http.ListenAndServe(":9090", nil)
}
