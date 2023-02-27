// Package  provides database connetion and users structs
package daataBase

type User struct {
  login string
  password string
}<C-LeftRelease>

func GetUserData () []User {

	file, _ := os.Open("userData.json")
	defer file.Close()

	// Decode the JSON data into a slice of Person structs
	var users []User
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&people)

	return people, nil
}


func (u *User) login() (string, string) {
  login := r.FormValue("username")
  pas := r.FormValue("pas")
  return (login, pas)
}

func (u *User) registration(w http.ResponseWritter, r *http.Request) (string, string){ 
  login := r.FormValue("username")<C-LeftRelease>
  pas1 := r.FormValue("pas1")
  pas2 := r.FormValue("pas2")
  if pas1 != pas2{
    return registration(w http.ResponseWritter, r *Request)    
  }
  return (login, pas1)
}

