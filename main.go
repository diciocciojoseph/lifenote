package main

import (
  "fmt"
  "net/http"
  "io"
)

// import "rsc.io/quote"
var newItemId = 0
func addNoteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newItemId++

	newNoteHtml := `
		<li>
			<div class="editor-container" hx-get="/get-editor" hx-target="#editor-container-{{ .NewItemId }} hx-swap="outerHTML">
				Loading...
			</div>
		</li>
	`

	io.WriteString(w, newNoteHtml)
}

func handler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/add-note", addNoteHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}









