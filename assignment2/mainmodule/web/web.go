package web

import (
  "net/http"
  // "html/template"
  "fmt"
)

func loginPage(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "LoginPage.html")
  if r.method = "POST" {
    login := r.FormValue("login")
    pas := r.FormValue("pas")
    if CheckLogin(login, pas) {
      
    }
  }
}

func registerPage(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "RegisterPage.html")
  if r.method = "POST" {
    login := r.FormValue("login")
    pas1 := r.FormValue("pas1")
    pas2 := r.FormValue("pas2")
    if CheckRegister(login, pas1, pas2){
      AddToDb(login, pas1)
    }
  }
} 

func mainPage(w http.ResponseWriter, r *http.Request){
  http.ServeFile("mainPage.html")
  // a lot of hardcode
}

func ConnectToServer() {
  http.HandleFunc("/login", loginPage)
  http.HandleFunc("/register", registerPage)
  http.HandleFunc("/", mainPage)
  fmt.Println("Server running")
  err := http.ListenAndServe("8181", nil)
  if err != nil{
    return
  }
}
