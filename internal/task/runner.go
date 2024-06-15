package task

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func TaskRunner(tasks []Task) {
	taskDone := make(map[string]chan struct{})
	var wg sync.WaitGroup

	// Initialize channels
	for _, task := range tasks {
		taskDone[task.Name] = make(chan struct{})
	}

	// Execute tasks
	for _, task := range tasks {
		wg.Add(1)
		go executeTask(task, taskDone, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()
}

func executeTask(task Task, taskDone map[string]chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Wait for dependencies to complete
	for _, dependency := range task.DependsOn {
		<-taskDone[dependency]
	}

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Printf("Started %v\n", task.Name)

	// Execute the task command
	cmd := exec.Command(task.Command[0], task.Command[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Printf("Error running %v: %v\n", task.Name, err)
	} else {
		fmt.Println(string(out))
		logger.Printf("%v done.\n", task.Name)
	}

	// Perform cleanup if needed
	if task.Cleanup {
		logger.Printf("Cleaning up %v\n", task.Name)
		cleanupCmd := exec.Command("rm", "-rf", task.CleanupPath)
		cleanupOut, cleanupErr := cleanupCmd.CombinedOutput()
		if cleanupErr != nil {
			logger.Printf("Error cleaning up %v: %v\n", task.Name, cleanupErr)
		} else {
			fmt.Println(string(cleanupOut))
			logger.Printf("Cleanup done for %v\n", task.Name)
		}
	}

	// Signal task completion
	close(taskDone[task.Name])
}