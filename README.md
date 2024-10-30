[# Task Manager CLI in Go

A simple command-line task manager built with Go. This project allows you to create, update, delete, and manage tasks saved in a JSON file. It supports various task statuses such as `todo`, `in-progress`, and `done`.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- Add, update, and delete tasks.
- Display tasks by status (`todo`, `in-progress`, `done`).
- Persist data in a JSON file.
- Simple CLI interface for quick task management.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/SamyDnx/task-manager-cli.git
   cd task-manager-cli
   ```

2. **Build the project**:
   Make sure you have Go installed. Run the following command in the project directory to build the CLI tool:
   ```bash
   go build -o taskmanager
   ```

3. **Set up the tasks file**:
   The program uses `tasks.json` for storing tasks. This file will be created automatically when you first run the program if it does not already exist.

## Usage

Run the program with the following syntax:

```bash
./taskmanager <command> [arguments]
```

For a list of available commands, use:
```bash
./taskmanager help
```

## Commands

| Command                  | Description                                                                                         |
|--------------------------|-----------------------------------------------------------------------------------------------------|
| `list`                   | Displays all tasks.                                                                                 |
| `list todo`              | Displays all tasks with status `todo`.                                                              |
| `list in-progress`       | Displays all tasks with status `in-progress`.                                                       |
| `list done`              | Displays all tasks with status `done`.                                                              |
| `add <description>`      | Adds a new task with the provided description and sets status to `todo`.                            |
| `update <id> <desc>`     | Updates the task with the specified ID to a new description.                                        |
| `delete <id>`            | Deletes the task with the specified ID.                                                             |
| `mark-todo <id>`         | Changes the status of the specified task to `todo`.                                                 |
| `mark-in-progress <id>`  | Changes the status of the specified task to `in-progress`.                                          |
| `mark-done <id>`         | Changes the status of the specified task to `done`.                                                 |

## Examples

### Add a New Task
```bash
./taskmanager add "Finish project report"
```

### List All Tasks
```bash
./taskmanager list
```

### Update a Task
```bash
./taskmanager update 1 "Finish and submit project report"
```

### Delete a Task
```bash
./taskmanager delete 1
```

### Change Status of a Task
```bash
./taskmanager mark-done 2
```

## License

This project is open-source and available under the MIT License. See the `LICENSE` file for more information.
](https://roadmap.sh/projects/task-tracker)
