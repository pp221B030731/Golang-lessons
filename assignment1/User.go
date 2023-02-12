package main

import (
  "fmt"
  "os"
  "encoding/csv"
)

type User struct {
  login string
  password string
}

func (u *User) addToData(dataFile *os.File) {
    writer := csv.NewWriter(dataFile)
    writer.Write([]string{u.login, u.password})
    writer.Flush()
}

func (u *User) checkLoginnRepaet(dataFile *os.File, login string) bool{
    reader := csv.NewReader(dataFile)
    data, _ := reader.ReadAll()
    for _, record := range data{
      if record[0] == login {
        return true
      }
    }
    return false
}


func (u *User) registration(dataFile *os.File) {
    var pas1 string
    var pas2 string
    var log string
    fmt.Println("Enter a login: ")
    fmt.Scanln(&log)
    fmt.Print("\033[H\033[2J")
    for u.checkLoginnRepaet(dataFile, log) {
      fmt.Println("There already user with same login! \n Enter a login: ")
      fmt.Scanln(&log)
      fmt.Print("\033[H\033[2J")
    }
    u.login = log
    for {
      fmt.Println("Enter a password: ")
      fmt.Scanln(&pas1)
      fmt.Println("Reapet Password: ")
      fmt.Scanln(&pas2)
      if pas1 == pas2{ 
        u.password = pas1
        break
      }
      fmt.Print("\033[H\033[2J")
      fmt.Println("Paswword is not same!")
    }
  fmt.Print("\033[H\033[2J")
  u.addToData(dataFile)
  fmt.Println("Registration Complet :)")
}


func (u *User) logining(dataFile *os.File) bool{
  var log, pas string
  fmt.Println("Enter a login: ")
  fmt.Scanln(&log)
  fmt.Println("Enter a password: ")
  fmt.Scanln(&pas)
  if CheckPassword(dataFile, log, pas) {
    u.login = log
    return true
  }
  return false
}

func CheckPassword(dataFile *os.File, log, pas string) bool{
  reader := csv.NewReader(dataFile)
  data, _ := reader.ReadAll()
  for _, record := range data {
    if record[0] == log {
      if record[1] == pas {
        return true
      } 
      return false
    }
  }
  return false
}
