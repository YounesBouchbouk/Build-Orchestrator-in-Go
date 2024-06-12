package manager

import (
	"fmt"

	"github.com/Younesbouchbouk/orchestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/*
   Accept requests from users to start and stop tasks
   Schedule tasks onto worker machines
   Keep track of tasks, their states, and the machine on which they run
*/

type Manager struct {
	Pending       queue.Queue                  // queue of pending taskes
	TaskDb        map[string][]*task.Task      //  list of taskes
	EventDb       map[string][]*task.TaskEvent // list of taskEvents
	Workers       []string                     // list of workers by name
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
