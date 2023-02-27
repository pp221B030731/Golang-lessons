package main

import (
  "mainmodule/web"
  "fmt"
  "mainmodule/daataBase"
  "sort"
)


func main() {
  // Data
  var userData []daataBase.User
  var productDataByP []daataBase.Product
  var productDataByR []daataBase.Product

  productDataByP = daataBase.GetProductData()
  productDataByR = productDataByP
  sort.Slice(productDataByR, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
  userData = daataBase.GetUserData()


  web.ConnectToServer()
}


