package main

import (
  "fmt"
  "net/http"
)

// import "rsc.io/quote"


func handler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func main(){
	http.HandleFunc("/", handler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}









