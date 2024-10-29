package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
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

	if argc == 0 {
		fmt.Println("Missing requiered argument. Run with help command to display all available commands.")
		return
	} else if argc == 1 && argv[0] == "help" {
		printHelp()
	} else if argc >= 1 && argv[0] == "list" {
		if argc == 1 {
			printList(tasks)
		} else if argv[1] == "todo" {
			displayStatusTasks(tasks, "todo")
		} else if argv[1] == "in-progress" {
			displayStatusTasks(tasks, "in-progress")
		} else if argv[1] == "done" {
			displayStatusTasks(tasks, "done")
		} else {
			fmt.Println("Unvalid command:", argv)
			return
		}
	} else if argc >= 1 && argv[0] == "add" {
		if argc == 1 {
			fmt.Println("Run <add [task you want to add]> to create a new task")
		} else if argc == 2 {
			addTask(tasks, argv[1], filename)
		} else {
			fmt.Println("Invalid command:", argv)
		}
	} else if argc >= 1 && argv[0] == "update" {
		if argc == 1 {
			fmt.Println("Run <update [id] [updated task description]> to update a task")
		} else if argc == 3 {
			updateTask(tasks, argv[1], argv[2], filename)
		} else {
			fmt.Println("Unvalid command:", argv)
		}
	} else if argc >= 1 && argv[0] == "delete" {
		if argc == 1 {
			fmt.Println("Run <delete [id]> to delete a task.")
		} else if argc == 2 {
			deleteTask(tasks, argv[1], filename)
		} else {
			fmt.Println("Invalid command:", argv)
			return
		}
	} else if argc >= 1 && argv[0] == "mark-done" {
		markStatusTask(tasks, argv[1], argv[0], filename)
	} else if argc >= 1 && argv[0] == "mark-in-progress" {
		markStatusTask(tasks, argv[1], argv[0], filename)
	} else if argc >= 1 && argv[0] == "mark-todo" {
		markStatusTask(tasks, argv[1], argv[0], filename)
	} else {
		fmt.Println("Invalid command:", argv)
		return
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

func printHelp() {
	fmt.Println("Here is all available commands:")
	fmt.Println("\tlist: \tdisplay all tasks")
	fmt.Println("\tadd: \tadd a task")
	fmt.Println("\tmark: \tset a new status of a task")
	fmt.Println("\tupdate: change the description of a task")
	fmt.Println("\tdelete: delete a task")
	fmt.Println("\nRun a command without anymore argument to learn how to use it.")
	fmt.Println()
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

func displayStatusTasks(tasks Tasks, status string) {
	for _, task := range tasks.Tasks {
		if task.Status == status {
			fmt.Println("Task id:", task.Id)
			fmt.Println("Task description:", task.Description)
			fmt.Println("Task created at:", task.CreatedAt)
			fmt.Println("Task updated at:", task.UpdatedAt)
			fmt.Println("Task status:", task.Status)
			fmt.Println()
		}
	}
}

func updateJSON(tasks Tasks, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error encoding tasks to JSON:", err)
		return
	}
}

func addTask(tasks Tasks, description string, filename string) {
	for _, task := range tasks.Tasks {
		if task.Description == description {
			fmt.Println("Task", description, "all ready exist.")
			return
		}
	}

	newTask := Task{
		Id:          len(tasks.Tasks) + 1,
		Description: description,
		CreatedAt:   time.Now().Format("2006/01/02, 15:04"),
		UpdatedAt:   time.Now().Format("2006/01/02, 15:04"),
		Status:      "todo",
	}

	tasks.Tasks = append(tasks.Tasks, newTask)

	updateJSON(tasks, filename)
	fmt.Println("Task added:", description)
}

func updateTask(tasks Tasks, _id string, description string, filename string) {
	id, err := strconv.Atoi(_id)
	if err != nil {
		fmt.Println("Invalid id:", _id)
		return
	}

	for _, task := range tasks.Tasks {
		if task.Description == description {
			fmt.Println("Task", description, "allready exists")
			return
		}
	}

	foundTask := false
	for i, task := range tasks.Tasks {
		if task.Id == id {
			tasks.Tasks[i].Description = description
			tasks.Tasks[i].UpdatedAt = time.Now().Format("2006/01/02, 15:04")
			foundTask = true
			break
		}
	}

	if foundTask == false {
		fmt.Println("Task not found.")
		return
	}

	updateJSON(tasks, filename)
	fmt.Println("Task updated")
}

func markStatusTask(tasks Tasks, _id string, _status string, filename string) {
	id, err := strconv.Atoi(_id)
	if err != nil {
		fmt.Println("Invalid id:", _id)
		return
	}

	status := ""
	if _status == "mark-done" {
		status = "done"
	} else if _status == "mark-in-progress" {
		status = "in-progress"
	} else if _status == "mark-todo" {
		status = "todo"
	} else {
		fmt.Println("Invalid status:", _status)
		return
	}

	foundTask := false
	for i, task := range tasks.Tasks {
		if task.Id == id {
			if task.Status == status {
				fmt.Println("Task allready have", status, "status.")
				return
			}
			tasks.Tasks[i].Status = status
			tasks.Tasks[i].UpdatedAt = time.Now().Format("2006/01/02, 15:04")
			foundTask = true
			break
		}
	}

	if foundTask == false {
		fmt.Println("Task not found.")
		return
	}

	updateJSON(tasks, filename)
	fmt.Println("Status modified.")
}

func deleteTask(tasks Tasks, _id string, filename string) {
	id, err := strconv.Atoi((_id))
	if err != nil {
		fmt.Println("Invalid id:", _id)
		return
	}

	foundTask := false
	for i, task := range tasks.Tasks {
		if task.Id == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			foundTask = true
			break
		}
	}

	if foundTask == false {
		fmt.Println("Task not found.")
		return
	} else {
		for i, task := range tasks.Tasks {
			if task.Id != i+1 {
				tasks.Tasks[i].Id = i + 1
			}
		}
	}

	updateJSON(tasks, filename)
	fmt.Println("Task deleted.")
}
