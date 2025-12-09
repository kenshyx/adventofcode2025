package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenshyx/adventofcode2025/day1"
	"github.com/kenshyx/adventofcode2025/day2"
	"github.com/kenshyx/adventofcode2025/day3"
	"github.com/kenshyx/adventofcode2025/day4"
	"github.com/kenshyx/adventofcode2025/day5"
	"github.com/kenshyx/adventofcode2025/day6"
	"github.com/kenshyx/adventofcode2025/day7"
	"github.com/kenshyx/adventofcode2025/day8"
	"github.com/kenshyx/adventofcode2025/day9"
	"github.com/kenshyx/adventofcode2025/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sessionCookie := os.Getenv("ADVENT_SESSION")
	authSession := &http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	}
	client := &http.Client{}
	authenticatedR := utils.CreateClientWithAuth(authSession, client)
	fmt.Printf("Solutions for day one: %+v\n",
		day1.GetSolution(authenticatedR.GetPuzzle(2025, 1)))
	fmt.Printf("Solutions for day two: %+v\n",
		day2.GetSolution(authenticatedR.GetPuzzle(2025, 2)))
	fmt.Printf("Solutions for day three: %+v\n",
		day3.GetSolution(authenticatedR.GetPuzzle(2025, 3)))
	fmt.Printf("Solutions for day four: %+v\n",
		day4.GetSolution(authenticatedR.GetPuzzle(2025, 4)))
	fmt.Printf("Solutions for day five: %+v\n",
		day5.GetSolution(authenticatedR.GetPuzzle(2025, 5)))
	fmt.Printf("Solutions for day six: %+v\n",
		day6.GetSolution(authenticatedR.GetPuzzle(2025, 6)))
	fmt.Printf("Solutions for day seven: %+v\n",
		day7.GetSolution(authenticatedR.GetPuzzle(2025, 7)))
	fmt.Printf("Solutions for day eight: %+v\n",
		day8.GetSolution(authenticatedR.GetPuzzle(2025, 8)))
	fmt.Printf("Solutions for day nine: %+v\n",
		day9.GetSolution(authenticatedR.GetPuzzle(2025, 9)))
}
