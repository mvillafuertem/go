package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("https://api.github.com/users/mvillafuertem")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Struct
	var model Model
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(model)
	fmt.Println("USERNAME ~> ", model.Username)
	fmt.Println("ID ~> ", model.ID)

	// Json
	// var apiResponse map[string]interface{}
	// err = json.NewDecoder(readCloser).Decode(&apiResponse)
	// fmt.Println(apiResponse)

	// var username string
	// username = apiResponse["login"].(string)
	// fmt.Println("USERNAME ~> ", username)
	// id := apiResponse["id"].(float64)
	// fmt.Println("ID ~> ", id)

	// StdOut
	// _, _ = io.Copy(os.Stdout, readCloser)
}

type Model struct {
	ID        int
	Username  string
	Bio       string
	Followers int
}

func (g *Model) UnmarshalJSON(b []byte) error {
	var gitHubResponse struct {
		ID        float64 `json:"id"`
		Username  string  `json:"login"`
		Bio       string  `json:"bio"`
		Followers float64 `json:"followers"`
	}

	if err := json.Unmarshal(b, &gitHubResponse); err != nil {
		return err
	}

	g.ID = int(gitHubResponse.ID)
	g.Username = gitHubResponse.Username
	g.Bio = gitHubResponse.Bio
	g.Followers = int(gitHubResponse.Followers)
	return nil
}
