package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrEmptyDesc    = errors.New("description cannot be empty")
)

type Storage struct {
	LastTaskID int     `json:"last_task_id"`
	Tasks      []*Task `json:"tasks"`
	path       string
}

func NewStorage(path string) *Storage {
	var file *os.File

	// check if tasks.json exists
	// if not create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if errF := createStorageFile(path); errF != nil {
			log.Fatalf("cannot create storage file: %v\n", errF)
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot open storage file: %v\n", err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read storage file: %v\n", err)
	}

	var s *Storage = &Storage{
		Tasks: make([]*Task, 0),
		path:  path,
	}
	if err := json.Unmarshal(buffer, s); err != nil {
		log.Fatalf("cannot unmarshal storage file: %v\n", err)
	}

	return s
}

func (s *Storage) AddTask(description string) error {
	task := &Task{
		ID:          s.LastTaskID + 1,
		Description: description,
		Status:      TaskStatusTODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.LastTaskID++
	s.Tasks = append(s.Tasks, task)

	return s.save()
}

func (s *Storage) UpdateTask(id int, description string) error {
	for _, task := range s.Tasks {
		if task.ID == id {
			task.Description = description
			task.UpdatedAt = time.Now()

			return s.save()
		}
	}

	return ErrTaskNotFound
}

func (s *Storage) DeleteTask(id int) error {
	for i, task := range s.Tasks {
		if task.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)

			return s.save()
		}
	}

	return ErrTaskNotFound
}

func (s *Storage) MarkInProgress(id int) error {
	for _, task := range s.Tasks {
		if task.ID == id {
			task.Status = TaskStatusInProgress
			task.UpdatedAt = time.Now()

			return s.save()
		}
	}

	return ErrTaskNotFound
}

func (s *Storage) MarkDone(id int) error {
	for _, task := range s.Tasks {
		if task.ID == id {
			task.Status = TaskStatusDone
			task.UpdatedAt = time.Now()

			return s.save()
		}
	}

	return ErrTaskNotFound
}

func (s *Storage) ListAllTasks() []*Task {
	return s.Tasks
}

func (s *Storage) ListTODOTasks() []*Task {
	tasks := make([]*Task, 0)

	for _, task := range s.Tasks {
		if task.Status == TaskStatusTODO {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func (s *Storage) ListDoneTasks() []*Task {
	tasks := make([]*Task, 0)

	for _, task := range s.Tasks {
		if task.Status == TaskStatusDone {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func (s *Storage) ListInProgressTasks() []*Task {
	tasks := make([]*Task, 0)

	for _, task := range s.Tasks {
		if task.Status == TaskStatusInProgress {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

// saves tasks in tasks.json
func (s *Storage) save() error {
	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, buffer, 0666)
}

// creates tasks.json with required fields
func createStorageFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(`{
	"last_task_id": 0,
	"tasks": []
}`))

	return err
}
