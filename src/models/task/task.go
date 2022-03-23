package task

type Task struct {
	text   string
	status bool
}

func CreateTask(text string) *Task {
	return &Task{text: text, status: false}
}

func (t *Task) GetText() string {
	return t.text
}

func (t *Task) SetText(text string) {
	t.text = text
}

func (t *Task) GetStatus() bool {
	return t.status
}

func (t *Task) SetStatus(status bool) {
	t.status = status
}
