package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Tasks struct {
	Tasks []Task `json:"Tasks"`
}

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Status      string `json:"status"`
}

func main() {
	argv := os.Args[1:]
	argc := len(argv)

	filename := "tasks.json"

	// create file if it does not exist
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		os.Create(filename)
	}

	tasks, err := readJSON(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if argc >= 1 && argv[0] == "list" {
		if argc == 1 {
			printList(tasks)
		} else if argv[1] == "todo" {
			// display all todo tasks
		} else if argv[1] == "in-progress" {
			// display all in-progress tasks
		} else if argv[1] == "done" {
			// display all done tasks
		} else {
			fmt.Println("Unvalid command:", argv)
		}
	}
}

func readJSON(filename string) (Tasks, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return Tasks{}, err
	}

	// data is of type []byte
	data, err := io.ReadAll(file)
	if err != nil {
		return Tasks{}, err
	}

	var tasks Tasks

	json.Unmarshal(data, &tasks)

	return tasks, nil
}

func printList(tasks Tasks) {
	for i := 0; i < len(tasks.Tasks); i++ {
		fmt.Println("Task id:", tasks.Tasks[i].Id)
		fmt.Println("Task description:", tasks.Tasks[i].Description)
		fmt.Println("Task created at:", tasks.Tasks[i].CreatedAt)
		fmt.Println("Task updated at:", tasks.Tasks[i].UpdatedAt)
		fmt.Println("Task status:", tasks.Tasks[i].Status)
		fmt.Println()
	}
}
