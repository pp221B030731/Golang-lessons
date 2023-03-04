// Package  provides ...
package daataBase

import (
  "sort"
)

type Session struct {
  UserDataList []User
  newUsersList []User
  ProductDataByP []Product
  ProductDataByR []Product
  ThisUser User
}

func (s *Session) Init(){
  s.ProductDataByP = GetProductData()
  s.ProductDataByR = GetProductData()
  sort.Slice(s.ProductDataByR, func(i, j int) bool {
		return s.ProductDataByR[i].Rating < s.ProductDataByR[j].Rating
  })
  sort.Slice(s.ProductDataByP, func(i, j int) bool {
		return s.ProductDataByP[i].Price < s.ProductDataByP[j].Price
  })

  s.UserDataList = GetUserData()
}

func (s *Session)CheckLogin(login string, pas string) bool{
  for _, i := range s.UserDataList {
    if i.Login == login {
      if pas == i.Password{
        s.ThisUser = i
        return true
      }
    }
  }
  return false
}


func (s *Session) Register(login string, pas string) {
  s.ThisUser = User{
    Login: login,
    Password: pas,
  }
  s.UserDataList = append(s.UserDataList, s.ThisUser)
  s.newUsersList = append(s.newUsersList, s.ThisUser)
}

