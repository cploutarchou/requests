package main

import (
	"fmt"
	"io"
	"time"

	"github.com/cploutarchou/go-http/http"
)

func getGithubClientWithOutConfig() http.Client {
	builder := http.NewBuilder()
	builder.SetRequestTimeout(50 * time.Second).SetResponseTimeout(50 * time.Second).SetMaxIdleConnections(10)
	http.Headers().SetHeaders(map[string][]string{
		"Content-Type": {"application/json"},
	}
	builder.SetHeaders().
	return builder.Build()
}

func getGithubClientWithConfig() http.Client {
	_client := http.NewBuilder()
	_client.SetConfig(&http.TimeoutSettings{
		MaxIdleConnections: 10,
		ResponseTimeout:    50 * time.Second,
		RequestTimeout:     50 * time.Second,
	})
	return _client
}

func getGithubClientBySetters() http.Client {
	_client := http.NewBuilder()
	_client.SetRequestTimeout(50 * time.Second)
	_client.SetResponseTimeout(50 * time.Second)
	_client.SetMaxIdleConnections(10)
	return _client
}
func init() {
	client = getGithubClientWithOutConfig()
	client = getGithubClientWithConfig()
	client = getGithubClientBySetters()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func main() {
	//GetExample()
	user := User{
		FirstName: "Christos",
		LastName:  "Ploutarchou",
		Username:  "username",
	}
	PostExample(user)
	GetExample()

}

func GetExample() {
	response, err := client.Get("https://api.github.com", nil)
	client.DisableTimeouts(true)
	client.MakeHeaders()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func PostExample(u User) {
	response, err := client.Post("https://api.github.com", nil, u)
	client.MakeHeaders()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//bytes, err := io.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(response.StatusCode)
}
