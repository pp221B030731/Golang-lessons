package daataBase

import (
  // import json reader
  "encoding/json"
  "os"
  "sort"
)

type Product struct {
  name string 'json:"name"'
  price int64 'json:"price"'
  rating float64 'json:"rating"'
  number int64 'json:"number"'
}

func GetProductData() []Product {
  file, _ := os.Open("productData.json")
  defer file.Close()

  var data []Product
  decoder := json.NewDecoder(file)
  _ := decoder.Decode(&data)

  sort.Slice(data, func(i, j int) bool {
		return people[i].price < people[j].name
	})
}



func Search(name string, list []Product){
  for _, p := range list {
    if p.name == name {
      fmt.Println("----------------------")
      fmt.Println("Name: ", p.name)
      fmt.Println("Price: ", p.price)
      fmt.Println("Rating: ", p.rating)
      fmt.Println("----------------------")
      return
    }
  }
  fmt.Println("Sorry, couldn't find your product")
}


func (p *Product) Rate (gRating float64) {
  if gRating > 5.0  { 
    gRating = 5
  } else if gRating < 0 {
    gRating = 0
  }
  p.rating = (p.rating*float64(p.number) + gRating)/float64(p.number+1)
  p.number++
}

func UpdateListPosition (list []Product, indToRate int) {
  var to int
  for ind := indToRate; ind < len(list)-1; ind++ {
    if list[ind+1].rating > list[indToRate].rating {
      to = ind
      break
    }
  }
  if to == indToRate {
    for ind := indToRate; ind > 0; ind-- {
      if list[indToRate].rating > list[ind-1].rating {
        to = ind
        break
      }
    }
  }
  p := list[indToRate]
  list = append(list[:indToRate], list[(indToRate+1):]...)
  list = append(list[:to], append([]Product{p}, list[to:]...)...)
}


