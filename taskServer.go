package main

import (
	"TaskExecutor/executor"
	"TaskExecutor/task"
	"TaskExecutor/taskQueue"
	"time"
)

func main() {
	//returns an instance of taskQueue
	taskQueue := taskQueue.NewTaskQueue()

	//adder loop to add 50 tasks to the task queue
	for i := 0; i < 50; i++ {
		newTask := (&task.Task{Id: i, CreationTime: time.Now(), TaskData: "send email"}).UpdateStatus(task.Untouched)
		newTask.UpdateIsComplete(false)
		taskQueue.Enqueue(newTask)
	}

	//executor goroutine
	//dequeues a task from task queue
	//executes a task only if its 'IsComplete' flag is false, else enqueues it back for the cleaner to clean the task.
	//enqueues the task after execution with updated status and IsCompleted flag
	go func() {
		for {
			t := taskQueue.Dequeue()
			if t != nil && !t.CheckIsComplete() {
				executor.RunTask(t)
			}
			taskQueue.Enqueue(t)
		}
	}()

	//cleaner goroutine that runs continuously and cleans the task queue by removing completed and timed out tasks
	//enqueues failed tasks with updated IsComplete flag
	go taskQueue.Clean()

	//to avoid the program from exiting
	select {}
}
