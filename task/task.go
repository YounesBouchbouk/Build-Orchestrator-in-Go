package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

/*
  User submits a task to the system. At this point, the task has been enqueued but is waiting to be scheduled.
  let’s call this initial state Pending. Once the system has figured out where to run the task, we can say it
  has been moved into a state of Scheduled. The Scheduled state means the system has determined there is a
  machine that can run the task, but it is in the process of sending the task to the selected machine, or
  the selected machine is in the process of starting the task. Next, if the selected machine successfully
  starts the task, it moves into the Running state. Upon a task completing its work successfully or being
  stopped by a user, the task moves into a state of Completed. If at any point the task crashes or stops
  working as expected, the task then moves into a state of Failed.
*/

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	Port          nat.PortSet
	PortBinding   map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
