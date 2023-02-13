package main

import (
  "fmt"
  "strconv"
)

type Product struct {
  name string
  price int64
  rating float64
  number int64
}

func findIndByPrice(list []Product, neededPri int64) int {
  l := 0
  r := len(list)-1
  if r < 0 {
    return 0
  }
  var m int
  for l <= r{
    m = l + (r-l)/2
    if list[m].price > neededPri {
      r = m-1
    } else if list[m].price < neededPri {
      l = m+1 
    } else {
      return m
    }
  }
  return l
}

func findIndByRating(list []Product, neededRate float64) int {
  l := 0
  r := len(list)-1
  if r < 0{
    return 0
  }
  var m int
  for l <= r{
    m = l + (r-l)/2
    if list[m].rating > neededRate {
      r = m-1
    } else if list[m].rating < neededRate {
      l = m+1 
    } else {
      return m
    }
  }
  return l
}

func sortBy(data [][]string, flag bool) []Product {
  list := []Product{}
  var p Product
  var ind int
  for _, record := range data { 
    p.name = record[0]
    p.price, _ = strconv.ParseInt(record[1], 10, 64)
    p.rating, _ = strconv.ParseFloat(record[2], 64)
    p.number, _ =strconv.ParseInt(record[3], 10, 64)
    if flag{
      ind = findIndByPrice(list, p.price)
    } else {
      ind = findIndByRating(list, p.rating)
    }
    if ind == 0 {
      list = append([]Product{p}, list...)
    } else if ind == len(list)-1 {
      list = append(list, p)
    } else {
      list = append(list[:ind], append([]Product{p}, list[ind:]...)...)
    }
  }
  return list
}

func search(name string, list []Product){
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

func showList(list []Product) {
  fmt.Println("index  |   Name    |    price    |     Rating")
  for i, record := range list {
    fmt.Println(i, record.name, record.price, record.rating)
  }
}

func (p *Product) rate (gRating float64) {
  if gRating > 5.0  { 
    gRating = 5
  } else if gRating < 0 {
    gRating = 0
  }
  p.rating = (p.rating*float64(p.number) + gRating)/float64(p.number+1)
  p.number++
}

func updateListPosition (list []Product, indToRate int) {
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


