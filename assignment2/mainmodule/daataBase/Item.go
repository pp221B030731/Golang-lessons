package daataBase

import (
  // import json reader
  "encoding/json"
  "os"
  "sort"
  "fmt"
)

type Product struct {
  Name string `json:"name"`
  Price int64 `json:"price"`
  Rating float64 `json:"rating"`
  Number int64 `json:"number"`
}

func GetProductData() []Product {
  file, _ := os.Open("daataBase/productData.json")
  defer file.Close()

  var data []Product
  decoder := json.NewDecoder(file)
  _ = decoder.Decode(&data)

  sort.Slice(data, func(i, j int) bool {
		return data[i].Price < data[j].Price
	})

  return data
}



func Search(name string, list []Product){
  for _, p := range list {
    if p.Name == name {
      fmt.Println("----------------------")
      fmt.Println("Name: ", p.Name)
      fmt.Println("Price: ", p.Price)
      fmt.Println("Rating: ", p.Rating)
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
  p.Rating = (p.Rating*float64(p.Number) + gRating)/float64(p.Number+1)
  p.Number++
}

func UpdateListPosition (list []Product, indToRate int) {
  var to int
  for ind := indToRate; ind < len(list)-1; ind++ {
    if list[ind+1].Rating > list[indToRate].Rating {
      to = ind
      break
    }
  }
  if to == indToRate {
    for ind := indToRate; ind > 0; ind-- {
      if list[indToRate].Rating > list[ind-1].Rating {
        to = ind
        break
      }
    }
  }
  p := list[indToRate]
  list = append(list[:indToRate], list[(indToRate+1):]...)
  list = append(list[:to], append([]Product{p}, list[to:]...)...)
}


