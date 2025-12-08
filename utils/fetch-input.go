package utils

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Solution struct {
	Part1, Part2 int
}

func FetchInput(url string) (*bufio.Reader, *http.Response) {

	if url == "" {
		return bufio.NewReader(os.Stdin), nil
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sessionCookie := os.Getenv("ADVENT_SESSION")
	client := &http.Client{}
	req, er := http.NewRequest("GET", url, nil)
	if er != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})
	resp, e := client.Do(req)
	if e != nil {
		panic(err)
	}

	reader := bufio.NewReader(resp.Body)

	return reader, resp
}
