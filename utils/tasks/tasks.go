package tasks

import (
	"fmt"
	"hash/fnv"

	uuid "github.com/satori/go.uuid"
)

// task function that do the works
type taskFunc func()

type TaskPool struct {
	tasks     []*task
	closeChan chan struct{}
}

// Job: add job to the pool
func (tp *TaskPool) Job(id uuid.UUID, fc func()) error {
	var hashCode uint32
	h := fnv.New32a()
	_, _ = h.Write(id.Bytes())
	hashCode = h.Sum32()
	// make sure that the msg from the same Id is allocate to the same routine to be processed in sequence
	return tp.tasks[hashCode&uint32(len(tp.tasks)-1)].job(taskFunc(fc))
}

func (t *task) job(fc taskFunc) error {
	select {
	case t.callbackChan <- fc:
		return nil
	default:
		return ErrNotFoundCallBack
	}
}

func newTaskPool(count int) *TaskPool {
	if count <= 0 {
		count = DefTaskPoolCount
	}

	taskPool := TaskPool{
		tasks:     make([]*task, count),
		closeChan: make(chan struct{}),
	}
	for i := range taskPool.tasks {
		taskPool.tasks[i] = makeTask(i, TaskSize, taskPool.closeChan)
		if taskPool.tasks[i] == nil {
			fmt.Printf("task %d is nil\n", i)
		}
	}
	return &taskPool
}

type task struct {
	index        int
	callbackChan chan taskFunc
	closeChan    chan struct{}
}

func (t *task) start() {
	for {
		select {
		case <-t.closeChan:

			return
		case fc := <-t.callbackChan:
			fc()
			// todo: add time management
		}
	}
}

func makeTask(index int, size int, close chan struct{}) *task {
	t := &task{
		index:        index,
		callbackChan: make(chan taskFunc, size),
		closeChan:    close,
	}
	go t.start()
	return t
}
