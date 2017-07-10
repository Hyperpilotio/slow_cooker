package scenario

import "errors"

type Scenario struct {
	Tasks  []string `json:"tasks"`
	Weight int      `json:"weight"`
}

type Task struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

type TaskContext struct {
	context map[string]interface{}
}

func NewTaskContext() *TaskContext {
	return &TaskContext{
		context: make(map[string]interface{}),
	}
}

func (context *TaskContext) Get(key string) (interface{}, error) {
	value, ok := context.Context[key]
	if !ok {
		return nil, errors.New("Unable to find key: " + key)
	}

	return value, nil
}

func (context *TaskContext) Set(key string, value interface{}) {
	context.Context[key] = value
}

func (context *TaskContext) Reset() {
	context.context = make(map[string]interface{})
}
