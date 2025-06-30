package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var timeLayout = "2006-01-02 15:04 GMT-07:00"

func searchIndex(idxTarget int) error {
	left, right := 1, arrT[len(arrT)-1].Id
	for left <= right {
		mid := (left + right) / 2
		if idxTarget > mid {
			left = mid + 1
		} else if idxTarget < mid {
			right = mid - 1
		} else {
			if arrT[mid-1].Status == "deleted" {
				return errors.New("task has been deleted")
			} else {
				return nil
			}
		}
	}
	return errors.New("index not found")
}

func Execute(args []string) error {
	if len(args) == 1 {
		return errors.New("command can't empty")
	}
	switch args[1] {
	case "add":
		if len(args) == 2 {
			return errors.New("add argument can't empty")
		} else if val := strings.Join(args[2:], " "); len(val) > 100 {
			return errors.New("task title too long")
		} else {
			return addTask(val)
		}
	case "update":
		if len(args) <= 3 {
			return errors.New("update argument to short or empty")
		} else if idx, err := strconv.Atoi(args[2]); err != nil || idx <= 0 {
			return errors.New("id must be a natural number")
		} else if err := searchIndex(idx); err != nil {
			return err
		} else {
			return updateTask(idx, strings.Join(args[2:], " "))
		}
	case "delete":
		if len(args) <= 2 {
			return errors.New("delete argument to short or empty")
		} else if len(args) > 3 {
			return errors.New("delete argument to lomg")
		} else if idx, err := strconv.Atoi(args[2]); err != nil {
			return errors.New("index must be a natural number")
		} else if err := searchIndex(idx); err != nil {
			return err
		} else {
			return deleteTask(idx)
		}
	case "list":
		switch len(args) {
		case 2:
			return listTask("all")
		case 3:
			return listTask(args[2])
		default:
			return errors.New("list argument to long")
		}
	default:
		if len(args) <= 2 {
			return errors.New("mark argument to short or empty")
		} else if len(args) > 3 {
			return errors.New("mark argument to lomg")
		} else if idx, err := strconv.Atoi(args[2]); err != nil {
			return errors.New("index must be a natural number")
		} else if err := searchIndex(idx); err != nil {
			return err
		} else if len(args[1]) > 4 && (args[1][5:] == "done" || args[1][5:] == "in-progress") {
			return markTask(idx, args[1][5:])
		}
		return errors.New("invalid command")
	}
}

func addTask(title string) error {
	newTask := taskS{
		Id:       len(arrT) + 1,
		Title:    title,
		Status:   "todo",
		CreateAt: time.Now().Format(timeLayout),
		UpdateAt: "",
	}
	arrT = append(arrT, newTask)
	return nil
}

func updateTask(idx int, title string) error {
	arrT[idx-1].Title = title
	arrT[idx-1].UpdateAt = time.Now().Format(timeLayout)
	return nil
}

func deleteTask(idx int) error {
	arrT[idx-1].Status = "deleted"
	return nil
}

func listTask(flag string) error {
	maxWidths := map[string]int{
		"Id":       len("Id"),
		"Task":     len("Task"),
		"CreateAt": len("CreateAt"),
		"UpdateAt": len("UpdateAt"),
	}

	var rows [][]string
	for _, val := range arrT {
		if (flag != "all" && val.Status != flag) || val.Status == "deleted" {
			continue
		}

		indicator := "[ ]"
		switch val.Status {
		case "done":
			indicator = "[✓]"
		case "in-progress":
			indicator = "[○]"
		}

		titleWithIndicator := indicator + val.Title
		rows = append(rows, []string{
			strconv.Itoa(val.Id),
			titleWithIndicator,
			val.CreateAt,
			val.UpdateAt,
		})

		if w := len(rows[len(rows)-1][0]); w > maxWidths["Id"] {
			maxWidths["Id"] = w
		}
		if w := len(rows[len(rows)-1][1]); w > maxWidths["Task"] {
			maxWidths["Task"] = w
		}
		if w := len(rows[len(rows)-1][2]); w > maxWidths["CreateAt"] {
			maxWidths["CreateAt"] = w
		}
		if w := len(rows[len(rows)-1][3]); w > maxWidths["UpdateAt"] {
			maxWidths["UpdateAt"] = w
		}
	}

	keys := []string{"Id", "Task", "CreateAt", "UpdateAt"}

	printBorder := func() {
		fmt.Print("+")
		for _, k := range keys {
			fmt.Print(strings.Repeat("-", maxWidths[k]+2) + "+")
		}
		fmt.Println()
	}

	printBorder()

	fmt.Print("|")
	for _, k := range keys {
		fmt.Printf(" %-*s |", maxWidths[k], k)
	}
	fmt.Println()

	printBorder()

	for _, row := range rows {
		fmt.Print("|")
		for i, k := range keys {
			fmt.Printf(" %-*s |", maxWidths[k], row[i])
		}
		fmt.Println()
	}

	printBorder()

	return nil
}

func markTask(idx int, val string) error {
	arrT[idx-1].Status = val
	return nil
}

func ShowHelp() {
	fmt.Println(strings.Repeat("-", 48))
	fmt.Println("Simple CLI Aplication for to do list daily tasks")
	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("\tUsage:")
	fmt.Println("  .\\todo.exe <command> <flags>")

	fmt.Println("\tAvailable Commands:")
	fmt.Println("  add <Title>")
	fmt.Println("  delete <Id>")
	fmt.Print("  update <Id> <NewTitle>\n\n")

	fmt.Println("  mark-in-progress <Id>")
	fmt.Print("  mark-done <Id>\n\n")

	fmt.Println("  list <flags (optional)>")
	fmt.Println("\tflags: <done>, <todo>, <in-progress>, <>")
}
