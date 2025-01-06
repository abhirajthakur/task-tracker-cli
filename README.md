# Task Manager CLI

A command-line task management application built with Go and Cobra. Easily manage your tasks with features for creating, updating, deleting, and filtering tasks through a simple command-line interface.

## Features

- **Task Management**

  - Create new tasks
  - Update existing tasks
  - Delete tasks
  - Mark tasks as "in progress" or "done"

- **Task Filtering**

  - View all tasks
  - Filter tasks by status (done/todo/in progress)
  - Simple and intuitive command structure

- **Data Persistence**
  - Tasks automatically saved to JSON file
  - Data persists between sessions

## Installation

1. Ensure you have Go installed on your system (version 1.16 or higher)
2. Clone the repository:

```bash
git clone https://github.com/abhirajthakur/task-tracker-cli
cd task-tracker-cli
```

3. Build the application:

```bash
go build
```

## Usage

### Adding a Task

```bash
./task-tracker-cli add "Buy groceries"
# Output: added new task: Buy groceries (ID: 1)
```

### Updating a Task

```bash
./task-tracker-cli update 1 "Buy groceries and cook dinner"
```

### Deleting a Task

```bash
./task-tracker-cli delete 1
```

### Marking Task Status

```bash
./task-tracker-cli mark-in-progress 1
./task-tracker-cli mark-done 1
```

### Listing Tasks

```bash
# List all tasks
./task-tracker-cli list

# List completed tasks
./task-tracker-cli list done

# List tasks in progress
./task-tracker-cli list in-progress

# List todo tasks
./task-tracker-cli list todo
```

## Command Reference

| Command            | Description               | Usage                                |
| ------------------ | ------------------------- | ------------------------------------ |
| `add`              | Creates a new task        | `add "Task Description"`             |
| `update`           | Updates an existing task  | `update <task_id> "New Description"` |
| `delete`           | Removes a task            | `delete <task_id>`                   |
| `mark-in-progress` | Marks task as in progress | `mark-in-progress <task_id>`         |
| `mark-done`        | Marks task as done        | `mark-done <task_id>`                |
| `list`             | Lists all tasks           | `list`                               |
| `list done`        | Lists completed tasks     | `list done`                          |
| `list todo`        | Lists pending tasks       | `list todo`                          |
| `list in-progress` | Lists tasks in progress   | `list in-progress`                   |

## Configuration

Tasks are stored in `task-tracker.json`. The file is automatically created when you add your first task.
