// Package  provides database connetion and users structs
package daataBase
import (
  // import json reader
  "encoding/json"
	"os"
	"io/ioutil"
)


type User struct {
  Login string `json:"login`
  Password string `json:"password"`
}

func GetUserData () []User {

	file, _ := os.Open("daataBase/userData.json")
	defer file.Close()

	// Decode the JSON data into a slice of Person structs
	var users []User
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&users)

	return users
}


func AddToUserData (u []User) {
	updatedData, _ := json.Marshal(u)
	_ = ioutil.WriteFile("daataBase/userData.json", updatedData, 0644)
}

