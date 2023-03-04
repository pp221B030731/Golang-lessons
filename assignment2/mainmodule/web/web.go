package web

import (
  "net/http"
  "html/template"
  "fmt"
  "mainmodule/daataBase"
  "github.com/gorilla/mux"
  "os/exec"
)

var s daataBase.Session

func loginPage(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    if s.CheckLogin(r.FormValue("login"), r.FormValue("pas")){
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  } else if r.Method == "Reg"{
    http.Redirect(w, r, "/register", http.StatusSeeOther)
  }
  http.ServeFile(w, r, "web/LoginPage.html")
}

func registerPage(w http.ResponseWriter, r *http.Request){
  if r.Method == http.MethodPost {
    login := r.FormValue("login")
    pas := r.FormValue("pas1")
    s.Register(login, pas)
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }

  http.ServeFile(w, r, "web/RegisterPage.html")
}

func mainPage(w http.ResponseWriter, r *http.Request){
  tmpl := template.Must(template.ParseFiles("web/mainPage.html"))
  err := tmpl.Execute(w, s.ProductDataByP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SortedByRating(w http.ResponseWriter, r *http.Request){
  tmpl := template.Must(template.ParseFiles("web/mainPage.html"))
  err := tmpl.Execute(w, s.ProductDataByR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func open(url string) error {
    var cmd string
    var args []string
    cmd = "xdg-open"
    args = append(args, url)
    return exec.Command(cmd, args...).Start()
}


func ConnectToServer() {
  s.Init()
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/login", loginPage) 
  myRouter.HandleFunc("/register", registerPage)
  myRouter.HandleFunc("/byRating", SortedByRating)
  myRouter.HandleFunc("/", mainPage)
  fmt.Println("Server running")
  open("http://localhost:8181/login")
  err := http.ListenAndServe(":8181", myRouter)
  if err != nil{
    fmt.Println(err)
    return
  }
}
