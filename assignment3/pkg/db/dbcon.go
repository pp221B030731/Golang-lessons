package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "strconv"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)


type Book struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Cost        uint   `json:"cost"`
}

type DB struct {
  DataBase *gorm.DB
}

func ConnectDb() (*DB, error){
  dsn := "host=postgres user=bookinguser password=admin dbname=booking port=5433 sslmode=disable"
  var db DB
  var err error
  db.DataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  db.DataBase.AutoMigrate(&Book{})
  return &db, err
}

func (db *DB) GetInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
  if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
  var bookToFind Book
  err = db.DataBase.First(&bookToFind, id).Error
  if err == gorm.ErrRecordNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		}
  json.NewEncoder(w).Encode(bookToFind)
}

func (db *DB) GetList(w http.ResponseWriter, r *http.Request) {
  var bookList []Book
  err := db.DataBase.Find(&bookList).Error
  if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
  json.NewEncoder(w).Encode(bookList)
}

func (db *DB) UpdateBook(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(mux.Vars(r)["id"])
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  var book Book

  err = db.DataBase.First(&book, id).Error
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  if r.Method == "GET" {
    http.ServeFile(w, r, "pkg/web/updateBook.html") 
  } else if r.Method == "POST"{
    r.ParseForm()

    book.Title = r.FormValue("title")
    book.Description = r.FormValue("description")
    
    err = db.DataBase.Save(&book).Error

    if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
    
    json.NewEncoder(w).Encode(book)
  }
}


func (db *DB) DeleteBook(w http.ResponseWriter, r *http.Request) {
  var bookToFind Book
  id, err := strconv.Atoi(mux.Vars(r)["id"])
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  err = db.DataBase.First(&bookToFind, id).Error
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  db.DataBase.Delete(&bookToFind)
	json.NewEncoder(w).Encode(bookToFind)
}

func (db *DB) GetInfoByTitle(w http.ResponseWriter, r *http.Request) {
  var bookToFind Book
  title := mux.Vars(r)["title"]
  err := db.DataBase.First(&bookToFind, "title = ?", title).Error
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  json.NewEncoder(w).Encode(bookToFind)
}

func (db *DB) AddBook(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    http.ServeFile(w, r, "pkg/web/addBook.html") 
  } else if r.Method == "POST"{
    r.ParseForm()
    var book Book
    book.Title = r.FormValue("title")
    book.Description = r.FormValue("description")
    cost, _ := strconv.ParseUint(r.FormValue("cost"), 10, 32)
    book.Cost = uint(cost)
    err := db.DataBase.Create(&book).Error
    if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
    
    json.NewEncoder(w).Encode(book)
  }
}

func (db *DB) GetBooksSorted(w http.ResponseWriter, r *http.Request) {
  var books []Book
  order := mux.Vars(r)["order"]//orrder : asc and desc
  err := db.DataBase.Order("cost " + order).Find(&books).Error
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return 
  }
  json.NewEncoder(w).Encode(books)
}
