package task

type Task struct {
	text   string
	status bool
}

func CreateTask(text string, status bool) *Task {
	return &Task{text: text, status: status}
}

func (t *Task) GetText() string {
	return t.text
}

func (t *Task) SetText(text string) {
	t.text = text
}

func (t *Task) SetStatus(status bool) {
	t.status = status
}
