package web

import (
  dcon"assignment3/pkg/db"
  "net/http"
  "log"
  "github.com/gorilla/mux"
	"os/exec"
)

func open(url string) error {
    var cmd string
    var args []string
    cmd = "xdg-open"
    args = append(args, url)
    return exec.Command(cmd, args...).Start()
}

func Connect() {
		
		db, err := dcon.ConnectDb()
		if err != nil{
			log.Fatal(err)
		}
		

    router := mux.NewRouter()

    router.HandleFunc("/books", db.GetList)
    router.HandleFunc("/book/{id}", db.GetInfo)
		router.HandleFunc("/books/add", db.AddBook)
		router.HandleFunc("/book/{id}/update", db.UpdateBook)
    router.HandleFunc("/book/{id}/delete", db.DeleteBook)
    router.HandleFunc("/books/sort/{order}", db.GetBooksSorted)
		router.HandleFunc("/book/title/{title}", db.GetInfoByTitle)

		/* 
		router.HandleFunc("/books", dcon.getBooks).Methods("GET")
    router.HandleFunc("/books/{id}", dcon.getBook).Methods("GET")
    router.HandleFunc("/books/{id}", dcon.updateBook).Methods("PUT")
		router.HandleFunc("/books", dcon.createBook).Methods("POST")
    router.HandleFunc("/books/{id}", dcon.deleteBook).Methods("DELETE")
    router.HandleFunc("/books/search", dcon.searchBooksByTitle).Methods("GET")
    router.HandleFunc("/books/sort/{order}", dcon.sortBooks).Methods("GET")
		*/
    
		open("http://localhost:8181/books")

		log.Fatal(http.ListenAndServe(":8181", router))

}
// Бля пиздец устал делать эту хуйню кароч над хтмлки сдлеать + как в прошлом ассайнменте и добавить Форм валю тип такого также пофиксит баги тут и бд вроде готово полностью над протетстить в конце ты делай а я по сьебам

