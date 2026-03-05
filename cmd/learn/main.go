package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"example.com/internal/learn"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" {
		usage()
		return
	}

	switch args[0] {
	case "list":
		printList()
	case "tree":
		printTree()
	case "topic":
		if len(args) < 2 {
			exitWithErr(errors.New("missing topic, usage: go run ./cmd/learn topic <topic-name>"))
		}
		printByTopic(args[1])
	case "run":
		if len(args) < 2 {
			exitWithErr(errors.New("missing demo id, usage: go run ./cmd/learn run <demo-id>"))
		}
		if err := runDemo(args[1]); err != nil {
			exitWithErr(err)
		}
	default:
		exitWithErr(fmt.Errorf("unknown command: %s", args[0]))
	}
}

func usage() {
	fmt.Println("Learn CLI")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  go run ./cmd/learn list          # list all demos")
	fmt.Println("  go run ./cmd/learn tree          # list by topic")
	fmt.Println("  go run ./cmd/learn topic <name>  # list one topic")
	fmt.Println("  go run ./cmd/learn run <demoID>  # run one demo")
}

func printList() {
	for _, d := range learn.List() {
		fmt.Printf("%-24s topic=%-12s level=%d  %s\n", d.ID, d.Topic, d.Level, d.Title)
	}
}

func printTree() {
	all := learn.List()
	for _, topic := range learn.Topics() {
		fmt.Printf("[%s]\n", topic)
		for _, d := range all {
			if d.Topic != topic {
				continue
			}
			fmt.Printf("  - %s (L%d): %s\n", d.ID, d.Level, d.Description)
		}
		fmt.Println()
	}
}

func printByTopic(topic string) {
	all := learn.List()
	found := false
	for _, d := range all {
		if d.Topic != topic {
			continue
		}
		found = true
		fmt.Printf("%-24s level=%d  %s\n", d.ID, d.Level, d.Description)
	}
	if !found {
		exitWithErr(fmt.Errorf("topic not found: %s", topic))
	}
}

func runDemo(id string) error {
	demo, ok := learn.Find(id)
	if !ok {
		return fmt.Errorf("demo not found: %s", id)
	}

	root, err := findModuleRoot()
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "run", demo.Path)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Printf("running %s (%s)\n\n", demo.ID, demo.Path)
	return cmd.Run()
}

func findModuleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("go.mod not found from current path")
		}
		dir = parent
	}
}

func exitWithErr(err error) {
	msg := strings.TrimSpace(err.Error())
	if msg != "" {
		fmt.Fprintln(os.Stderr, "error:", msg)
	}
	os.Exit(1)
}
