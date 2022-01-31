package executor

import (
	"TaskExecutor/task"
	"fmt"
	"math/rand"
	"time"
)

//runs the task and updates the status
//assumption- randomly updates the status based on divisibility by 3
func RunTask(t *task.Task) {

	if t == nil {
		return
	}
	//check for timeout
	if (time.Now().Sub(t.CreationTime).Seconds()) > 10 {
		t.UpdateStatus(task.Timeout)
	}
	//random failing logic
	if rand.Intn(100)%3 == 0 {
		t.UpdateStatus(task.Failed)
	} else {
		t.UpdateStatus(task.Completed)
	}
	t.UpdateIsComplete(true)
	fmt.Printf("executed ID: %d, status:%s\n", t.Id, t.Status())
}
