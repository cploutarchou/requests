package main

import (
	"fmt"
	"io"
	"time"

	"github.com/cploutarchou/requests/http"
)

var client http.Client

func getGithubClientWithOutConfig() http.Client {
	builder := http.NewBuilder()
	builder.SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}

func getGithubClientBySetters() http.Client {
	_client := http.NewBuilder()
	_client.
		SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	_client.Headers().
		SetAcceptCharset("utf-8").
		SetAccept("application/json")
	return _client.Build()
}

func init() {
	// client = getGithubClientWithOutConfig()
	client = getGithubClientBySetters()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func main() {
	user := User{
		FirstName: "Christos",
		LastName:  "Ploutarchou",
		Username:  "username",
	}
	PostExample(user)
	GetExample()
}

func GetExample() {
	client.DisableTimeouts()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	client.EnableTimeouts()
	response, err = client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	bytes, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func PostExample(u User) {
	response, err := client.Post("https://api.github.com", nil, u)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
