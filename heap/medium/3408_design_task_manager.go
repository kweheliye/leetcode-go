package medium

import (
	"container/heap"
)

type PriorityQueueTask []*Task

func (pq PriorityQueueTask) Len() int { return len(pq) }

func (pq PriorityQueueTask) Less(i, j int) bool {
	if pq[i].priority != pq[j].priority {
		return pq[i].priority > pq[j].priority // higher priority first
	}
	return pq[i].taskId > pq[j].taskId // larger taskId first
}

func (pq PriorityQueueTask) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueueTask) Push(x interface{}) {
	*pq = append(*pq, x.(*Task))
}

func (pq *PriorityQueueTask) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type Task struct {
	userId   int
	taskId   int
	priority int
}

type TaskManager struct {
	taskMap map[int]*Task
	pq      PriorityQueueTask
}

func Constructor(tasks [][]int) TaskManager {
	manager := TaskManager{
		taskMap: make(map[int]*Task),
		pq:      make(PriorityQueueTask, 0),
	}

	for _, task := range tasks {
		manager.Add(task[0], task[1], task[2])
	}
	heap.Init(&manager.pq)
	return manager
}

// Add a task
func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{userId: userId, taskId: taskId, priority: priority}
	this.taskMap[taskId] = task
	heap.Push(&this.pq, task)
}

// Edit a task
func (this *TaskManager) Edit(taskId int, newPriority int) {
	if old, ok := this.taskMap[taskId]; ok {
		newTask := &Task{userId: old.userId, taskId: taskId, priority: newPriority}
		this.taskMap[taskId] = newTask
		heap.Push(&this.pq, newTask) // old task will be skipped later
	}
}

// Rmv Remove a task
func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskMap, taskId) // lazy delete
}

func (this *TaskManager) ExecTop() int {
	for this.pq.Len() > 0 {
		top := this.pq[0]
		valid, ok := this.taskMap[top.taskId]
		if !ok || valid != top {
			heap.Pop(&this.pq) // outdated or removed
			continue
		}
		heap.Pop(&this.pq)
		delete(this.taskMap, valid.taskId)
		return valid.userId
	}
	return -1
}
