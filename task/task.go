package task

import (
	"sync"
	"time"
)

//constants that denote the types of statuses accepted by the status field of the task type
const (
	Failed    string = "failed"
	Completed string = "completed"
	Timeout   string = "timeout"
	Untouched string = "untouched"
)

type Task struct {
	Id           int
	isComplete   bool
	status       string       // untouched, completed, failed, timeout
	CreationTime time.Time    // when was the task created
	TaskData     string       // field containing data about the task
	lock         sync.RWMutex //to lock and unlock whenever needed
}

//updates the status field
func (t *Task) UpdateStatus(newStatus string) *Task {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.status = newStatus
	return t
}

//returns the value of the status field
func (t *Task) Status() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.status
}

//returns the value of 'IsCompleted' field
func (t *Task) CheckIsComplete() bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.isComplete
}

//updates 'IsCompleted' field
func (t *Task) UpdateIsComplete(flag bool) *Task {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.isComplete = flag
	return t
}
