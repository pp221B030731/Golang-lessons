package main

import (
  "os"
  "fmt"
  "encoding/csv"
)


func main() {
  // FIle opendnig and working with data
  var todo int
  var searchName string
  autorized := false
  var u User
  var indToRate int
  var rateToGive float64
  userDataFile, err := os.OpenFile("users.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  if err != nil{
    fmt.Println("Error in opening user data base file") 
  }
  defer userDataFile.Close()

  productDataFile, err := os.OpenFile("products.csv", os.O_APPEND, 0660)
  if err != nil{
    fmt.Println("Error in opening product data base file")
  }
  defer productDataFile.Close()
  reader := csv.NewReader(productDataFile)
  data, _ := reader.ReadAll()

  productsListByPrice := sortBy(data, true)
  productsListByRating := sortBy(data, false) 

  // Autorization menue

  for !autorized{  
    fmt.Println("Good day dear user: \n 1: Sign in \n 2: Log in \n 0: To log out")
    fmt.Scanln(&todo)
    if todo == 1 {
      fmt.Print("\033[H\033[2J")
      u.registration(userDataFile)
      return
    } else if todo == 2 {
      fmt.Print("\033[H\033[2J")
      for {
        token := u.logining(userDataFile)
        if token {
          fmt.Print("\033[H\033[2J")
          fmt.Println("Wlcome ", u.login)
          autorized = true
          break
        }
        fmt.Print("\033[H\033[2J")
        fmt.Println("Not correct login or password!")
      }     
    } else if todo == 0 {
      fmt.Print("\033[H\033[2J")
      fmt.Println("GOOD by <3")
      return
    }
  }

  // Main menue


  var todo2 int
  flag := true
  for {
    if flag{
      showList(productsListByPrice)
      fmt.Println("--------------------------------")
      flag = false
    }
    fmt.Println("To filter: 0 \nTo rate item: 1\nTo find by name: 2 \nTo log out: 9")
    fmt.Scanln(&todo)
    switch todo {
    case 0:
      fmt.Println("To filter:\n by price: 0 \n by rating: 1\nTo go back: 9")
      fmt.Scanln(&todo2)
      switch todo2 {
      case 0:
        fmt.Print("\033[H\033[2J")
        showList(productsListByPrice)
        fmt.Println("--------------------------------")
      case 1:  
        fmt.Print("\033[H\033[2J")
        showList(productsListByRating)
        fmt.Println("--------------------------------")
      case 9: 
        fmt.Print("\033[H\033[2J")
        continue  
      default:
        
      }
    case 1: 
        fmt.Print("\033[H\033[2J")
        showList(productsListByRating)
        fmt.Println("--------------------------------")
        fmt.Println("Write a index of a list above:")
        fmt.Scanln(&indToRate)
        fmt.Println("Write a rating you want to give (0.0 - 5.0): ")
        fmt.Scanln(&rateToGive)
        if indToRate > 5 {
          indToRate = 5
        }
        if indToRate < 0 {
          indToRate = 0
        }
        productsListByRating[indToRate].rate(rateToGive)
        updateListPosition(productsListByRating, indToRate)
        fmt.Print("\033[H\033[2J")
        showList(productsListByRating)
        fmt.Print("--------------------------------")
      case 2:
      fmt.Println("Write a name :")
      fmt.Scanln(&searchName)
      search(searchName, productsListByPrice)
    case 9: 
      fmt.Print("\033[H\033[2J")
      fmt.Println("Good by", u.login)
      return
    default:
      
    }
  }

}
