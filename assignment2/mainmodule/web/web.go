package web

import (
  "net/http"
  // "html/template"
  "fmt"
)

func mainPage(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "mainPage.html")
}

var a = 10

func connectToServer() {
  http.HandleFunc("/", mainPage)
  fmt.Println("Server running")
  err := http.ListenAndServe("8181", nil)
  if err != nil{
    return
  }
}
