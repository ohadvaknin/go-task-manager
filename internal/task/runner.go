package task

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func TaskRunner(tasks []Task) {
	var wg sync.WaitGroup
	taskDone := make(chan bool, len(tasks))
	for _, v := range tasks {
		wg.Add(1)
		go executeTask(v, taskDone, &wg)
	}
	wg.Wait()
	close(taskDone)
}

func executeTask(task Task,taskDone chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Printf("Started %v\n", task.Name)
	cmd := exec.Command(task.Command[0], task.Command[1:]...)
	out1, err1 := cmd.CombinedOutput()
	if err1 != nil {
		logger.Printf("Error running %v\n", task.Name)
	} else {
		fmt.Println(string(out1))
		logger.Printf("%v done.", task.Name)
	}
	if task.Cleanup {
		logger.Printf("Cleaning up %v\n", task.Name)
		cmd := exec.Command("rm", "-rf", task.CleanupPath)
		out2, err2 := cmd.CombinedOutput()
		if err2 != nil {
			logger.Printf("Error cleaning up %v\n", task.Name)
		} else {
			fmt.Println(string(out2))
			logger.Printf("Cleanup done for %v\n", task.Name)
		}
	}
	taskDone <- true

}
