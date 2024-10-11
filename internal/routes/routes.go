package routes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bulgil/task-cli/internal/storage"
)

const (
	emptyDescription = "task description cannot be empty"
	emptyID          = "task id cannot be empty"
	invalidID        = "task id must be an integer"
	invalidCommand   = "invalid command"
)

type Router struct {
	storage *storage.Storage
}

func NewRouter(storage *storage.Storage) *Router {
	return &Router{storage: storage}
}

func (r *Router) Route(input []string) {
	switch input[0] {
	case "add":
		if len(input) < 2 {
			fmt.Println(emptyDescription)
			return
		}

		r.storage.AddTask(strings.Join(input[1:], " "))

	case "delete":
		if len(input) < 2 {
			fmt.Println(emptyID)
			return
		}

		id, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println(invalidID)
			return
		}

		r.storage.DeleteTask(id)

	case "update":
		id, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println(invalidID)
			return
		}

		if len(input) < 3 {
			fmt.Println(emptyDescription)
			return
		}

		r.storage.UpdateTask(id, strings.Join(input[2:], " "))

	case "list":
		if len(input) < 2 {
			tasks := r.storage.ListAllTasks()
			for _, task := range tasks {
				fmt.Println(task.Description)
			}
			fmt.Println("list")
			return
		}

		switch input[1] {
		case "todo":
			tasks := r.storage.ListTODOTasks()
			for _, task := range tasks {
				fmt.Println(task.Description)
			}
			return

		case "done":
			tasks := r.storage.ListDoneTasks()
			for _, task := range tasks {
				fmt.Println(task.Description)
			}
			return

		case "in-progress":
			tasks := r.storage.ListInProgressTasks()
			for _, task := range tasks {
				fmt.Println(task.Description)
			}
			return
		default:
			fmt.Println(invalidCommand)
			return
		}
	case "mark-in-progress":
		id, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println(invalidID)
			return
		}

		r.storage.MarkInProgress(id)

	case "mark-done":
		id, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println(invalidID)
			return
		}

		r.storage.MarkDone(id)
	default:
		fmt.Println(invalidCommand)
	}
}
