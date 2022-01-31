package taskQueue

import (
	"TaskExecutor/task"
	"fmt"
	"sync"
)

type TaskQueue struct {
	queue     *[]*task.Task
	taskQLock sync.Mutex //mutex to lock and unlock as needed
}

//enqueues a task into the task queue
func (tq *TaskQueue) Enqueue(t *task.Task) {
	tq.taskQLock.Lock()
	defer tq.taskQLock.Unlock()
	*tq.queue = append(*tq.queue, t)
}

//removes a task from the beginning of the task queue
func (tq *TaskQueue) Dequeue() *task.Task {
	return tq.RemoveFromQueue(0)

}

//initializes a new task queue
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{queue: &[]*task.Task{}}
}

//cleans the task queue
func (tq *TaskQueue) Clean() {
	for len(*tq.queue) != 0 {
		for idx, val := range *tq.queue {
			if val.CheckIsComplete() == true {
				if val.Status() == task.Timeout {
					fmt.Printf("ID: %d timed out, removing it from the queue\n", val.Id)
					tq.RemoveFromQueue(idx)
				} else if val.Status() == task.Completed {
					fmt.Printf("ID: %d is complete, removing it from the queue\n", val.Id)
					tq.RemoveFromQueue(idx)
				} else if val.Status() == task.Failed {
					fmt.Printf("ID: %d failed, adding it back to queue\n", val.Id)
					tq.RemoveFromQueue(idx)
					val.UpdateIsComplete(false)
					tq.Enqueue(val)
				}
			}
		}
	}
}

//used to remove a task from a particular index from the queue
func (tq *TaskQueue) RemoveFromQueue(idx int) *task.Task {
	var temp *task.Task
	tq.taskQLock.Lock()
	defer tq.taskQLock.Unlock()
	if len(*tq.queue) > idx {
		temp = (*tq.queue)[idx]
		*tq.queue = append((*tq.queue)[0:idx], (*tq.queue)[idx+1:]...)
	}
	return temp
}
