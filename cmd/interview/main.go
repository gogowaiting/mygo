package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"example.com/internal/interview"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		list()
		return
	}
	if args[0] == "help" {
		usage()
		return
	}

	switch args[0] {
	case "list":
		list()
	case "topics":
		topics()
	case "show":
		if len(args) < 2 {
			exitWithErr("usage: go run ./cmd/interview show <question-id>")
		}
		show(args[1])
	case "random":
		randomQuestion()
	default:
		exitWithErr("unknown command: " + args[0])
	}
}

func usage() {
	fmt.Println("Interview CLI")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  go run ./cmd/interview list")
	fmt.Println("  go run ./cmd/interview topics")
	fmt.Println("  go run ./cmd/interview show <question-id>")
	fmt.Println("  go run ./cmd/interview random")
}

func list() {
	for _, q := range interview.List() {
		fmt.Printf("%-20s topic=%-12s level=%d  %s\n", q.ID, q.Topic, q.Level, q.Prompt)
	}
}

func topics() {
	for _, t := range interview.Topics() {
		fmt.Println(t)
	}
}

func show(id string) {
	q, ok := interview.Find(id)
	if !ok {
		exitWithErr("question not found: " + id)
	}
	printQuestion(q)
}

func randomQuestion() {
	all := interview.List()
	if len(all) == 0 {
		exitWithErr("no questions")
	}
	rand.Seed(time.Now().UnixNano())
	q := all[rand.Intn(len(all))]
	printQuestion(q)
}

func printQuestion(q interview.Question) {
	fmt.Printf("ID: %s\n", q.ID)
	fmt.Printf("Topic: %s\n", q.Topic)
	fmt.Printf("Level: %d\n", q.Level)
	fmt.Printf("Q: %s\n", q.Prompt)
	fmt.Printf("A: %s\n", q.Answer)
	if len(q.Keywords) > 0 {
		fmt.Printf("Keywords: %s\n", strings.Join(q.Keywords, ", "))
	}
}

func exitWithErr(msg string) {
	msg = strings.TrimSpace(msg)
	if msg != "" {
		fmt.Fprintln(os.Stderr, "error:", msg)
	}
	os.Exit(1)
}
