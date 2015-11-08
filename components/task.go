package components

type Task struct {
	Id      int
	Title   string
	PhaseId int
}

type TaskPerson struct {
	Role   string
	UserId int
}
