package main

import (
  "fmt"
  "net/http"
)

// import "rsc.io/quote"


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, World!")
}

func main(){
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}









