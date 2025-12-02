package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

func (n *Node[T]) MoveNext() *Node[T] {
	if n == nil {
		return nil
	}
	return n.Next
}

func (n *Node[T]) MovePrev() *Node[T] {
	if n == nil {
		return nil
	}
	return n.Prev
}

func NewCircularList(n int) *Node[int] {
	if n <= 0 {
		return nil
	}
	head := &Node[int]{Value: 0}
	prev := head
	for i := 1; i < n; i++ {
		node := &Node[int]{Value: i}
		prev.Next = node
		node.Prev = prev
		prev = node
	}
	head.Prev = prev
	prev.Next = head

	return head
}

var initialPosition = 50
var moveRight = "R"
var moveLeft = "L"

func DayOne() int {
	safeDial := NewCircularList(100)
	password := 0
	currentNode := safeDial
	for i := 0; i < initialPosition; i++ {
		currentNode = currentNode.MoveNext()
	}
	file, _ := os.Open("day1-input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := bufio.NewReader(file)
	re := regexp.MustCompile(`\d+`)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		moves := re.FindString(line)
		m, _ := strconv.Atoi(moves)
		if currentNode == nil {
			continue
		}
		if strings.HasPrefix(line, moveLeft) {
			for l := 0; l < m; l++ {
				if currentNode.Value == 0 && l != 0 {
					password++
				}
				currentNode = currentNode.MovePrev()
			}
		}

		if strings.HasPrefix(line, moveRight) {
			for r := 0; r < m; r++ {
				if currentNode.Value == 0 && r != 0 {
					password++
				}
				currentNode = currentNode.MoveNext()
			}
		}

		if currentNode != nil && currentNode.Value == 0 {
			password++
		}

	}
	return password
}
