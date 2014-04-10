package batchqueue

type Runner struct {
	time    uint64  //任务的最早开始执行时间
	timeout uint64  //任务的执行超时时间
	task    Task    //任务
	next    *Runner //下一个对象
}

type Task interface {
	BatchRun(queue *Queue, tasks []Task)
}

//该对象尚未使用
type TaskNode struct {
	task *Runner
	next *TaskNode
}

//任务列表
type TaskList struct {
	head *Runner
}
