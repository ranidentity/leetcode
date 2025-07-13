package proto

type Task struct {
	ID      int
	Payload string
}

type FanInTask struct {
	ID        string
	DependsOn string // ID of another task this task depends on
	Payload   string
}
