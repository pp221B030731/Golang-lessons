package pkg 

import (
	"encoding/json"
	"net/http"
	env"assignment4"
	"fmt"
	    "io/ioutil"
)


type UnsplashResponse struct {
		ID      string `json:"id"`
		Urls    struct {
			Regular string `json:"regular"`
		} `json:"urls"`
}


func RandomPhoto() string {
	url := "https://api.unsplash.com/photos/random?client_id=" + env.UNSPLASH_TOKEN
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ures UnsplashResponse
	err = json.Unmarshal(body, &ures)
	if err !=nil{
		fmt.Println(err.Error)
	}
	photo := ures.Urls.Regular
	return photo
}
