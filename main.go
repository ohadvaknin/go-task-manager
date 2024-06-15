package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ohadvaknin/go-task-manager/internal/task"
)

func main() {
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		log.Printf("could not open file: %v\n", err)
		return
	}
	defer jsonFile.Close()
	var tasks []task.Task
	err = json.NewDecoder(jsonFile).Decode(&tasks)
	if err != nil {
		log.Printf("could not decode file: %v\n", err)
		return
	}
	task.TaskRunner(tasks)
}