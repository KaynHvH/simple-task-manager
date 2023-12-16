package main

import (
	"fmt"
	"strconv"
)

type Tasks struct {
	ID   int
	Name string
	Done bool
}

var taskList []Tasks

func (t *Tasks) addTask() {
	var newTaskName string
	fmt.Println("What's the name of your task?")
	_, err := fmt.Scanln(&newTaskName)
	if err != nil {
		fmt.Println(err)
	}

	newTask := Tasks{
		ID:   len(taskList) + 1,
		Name: newTaskName,
		Done: false,
	}

	taskList = append(taskList, newTask)
	fmt.Printf("Task '%s' added to the list.\n", newTaskName)
}

func (t *Tasks) setDone() {
	var taskIDString string
	fmt.Println("Enter the ID of the task that has been completed")
	_, err := fmt.Scanln(&taskIDString)
	if err != nil {
		fmt.Println("You have to enter task ID")
	}

	taskID, err := strconv.Atoi(taskIDString)
	if err != nil {
		fmt.Println("Invalid task ID format")
		return
	}

	for i := range taskList {
		if taskList[i].ID == taskID {
			taskList[i].Done = true
			fmt.Printf("Task with ID %d has been marked as done\n", taskID)
			return
		}
	}
	fmt.Printf("Task not found with ID: %d", taskID)
}

func (t *Tasks) deleteTasks() {
	var taskIDString string
	fmt.Println("Enter the ID of the task that you want to delete")
	_, err := fmt.Scanln(&taskIDString)
	if err != nil {
		fmt.Println("You have to enter task ID")
	}

	taskID, err := strconv.Atoi(taskIDString)
	if err != nil {
		fmt.Println("Invalid task ID format")
		return
	}

	for i := range taskList {
		if taskList[i].ID == taskID {
			taskList = append(taskList[:i], taskList[i+1:]...)
			fmt.Printf("Task with ID %d has been deleted\n", taskID)

			for j := i; j < len(taskList); j++ {
				taskList[j].ID = j + 1
			}
			return
		}
	}
	fmt.Println("Task not found with ID:", taskID)
}

func (t *Tasks) showTasks() {
	for _, i := range taskList {
		var isDone string
		if i.Done == false {
			isDone = "No"
		} else {
			isDone = "Yes"
		}

		fmt.Printf("Task: %s | Done: %s | ID: %d\n", i.Name, isDone, i.ID)
	}
}

func main() {
	for {
		var taskInstance Tasks

		fmt.Println("Enter your choice:\n" +
			"1. Add task\n" +
			"2. Set task done\n" +
			"3. Delete task\n" +
			"4. Show tasks")

		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
		}

		if choice == "1" {
			taskInstance.addTask()
		} else if choice == "2" {
			taskInstance.setDone()
		} else if choice == "3" {
			taskInstance.deleteTasks()
		} else if choice == "4" {
			taskInstance.showTasks()
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
