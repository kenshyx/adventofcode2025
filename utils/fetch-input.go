package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type Solution struct {
	Part1, Part2 int
}

type UrlWithAuth struct {
	url    string
	auth   *http.Cookie
	client *http.Client
}

func CreateClientWithAuth(auth *http.Cookie, client *http.Client) *UrlWithAuth {
	return &UrlWithAuth{auth: auth, client: client}
}

const UrlTpl = "https://adventofcode.com/%d/day/%d/input"

func (resource *UrlWithAuth) GetPuzzle(year int, day int) *UrlWithAuth {
	resource.url = fmt.Sprintf(UrlTpl, year, day)
	return resource
}

func (resource *UrlWithAuth) FetchInput() (*bufio.Reader, *http.Response) {

	if resource.url == "" {
		return bufio.NewReader(os.Stdin), nil
	}

	req, er := http.NewRequest("GET", resource.url, nil)
	if er != nil {
		panic(er)
	}
	req.AddCookie(resource.auth)
	resp, e := resource.client.Do(req)
	if e != nil {
		panic(e)
	}

	reader := bufio.NewReader(resp.Body)

	return reader, resp
}
