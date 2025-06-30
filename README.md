# Task Tracker CLI

A simple and efficient command-line tool built with Go to manage your daily tasks. Never miss a deadline again!

## ✨ Features

* **Add, Update, and Delete Tasks**: Easily manage your to-do list from the command line.
* **List Tasks**: View all your tasks, or filter by status (`todo`, `in-progress`, `done`).
* **Mark Tasks**: Update the status of your tasks to keep track of your progress.
* **Persistent Storage**: Your tasks are saved to a `Task.json` file, so you'll never lose your data.

## 💡 Inspiration

This project idea is taken from the list of projects on [roadmap.sh](https://roadmap.sh/projects). You can find the original project brief here: [Task Tracker Project](https://roadmap.sh/projects/task-tracker).

## ⚙️ Prerequisites

Before you begin, ensure you have the following installed on your system:

* **Go**: Version 1.16 or higher. You can download it from [go.dev](https://go.dev/dl/).

## 🚀 How To Use

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/FrenaldyH/Task-Tracker.git](https://github.com/FrenaldyH/Task-Tracker.git)
    cd Task-Tracker
    ```

2.  **Build the project:**
    ```bash
    go build .
    ```

3.  **Run the application:**
    You can now use the `todo` executable to manage your tasks. Here are the available commands:

    * **Add a new task:**
        ```bash
        ./todo add "My new awesome task"
        ```

    * **List all tasks:**
        ```bash
        ./todo list
        ```

    * **List tasks with a specific status:**
        ```bash
        ./todo list todo
        ./todo list in-progress
        ./todo list done
        ```

    * **Update a task:**
        ```bash
        ./todo update <task_id> "My updated task title"
        ```

    * **Mark a task as in-progress:**
        ```bash
        ./todo mark-in-progress <task_id>
        ```

    * **Mark a task as done:**
        ```bash
        ./todo mark-done <task_id>
        ```

    * **Delete a task:**
        ```bash
        ./todo delete <task_id>
        ```

    * **Show help:**
        ```bash
        ./todo
        ```

## 📜 Commands

Here's a summary of all the available commands:

| Command | Description |
| --- | --- |
| `add <title>` | Adds a new task with the given title. |
| `update <id> <new_title>` | Updates the title of a task. |
| `delete <id>` | Deletes a task. |
| `list [status]` | Lists all tasks or filters by status. |
| `mark-in-progress <id>` | Marks a task as "in-progress". |
| `mark-done <id>` | Marks a task as "done". |

---

<p align="center">
  Developed with by FrenaldyH
</p>
