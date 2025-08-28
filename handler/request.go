package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type%s) is required", name, typ)
}

// CreateTodo
type CreateTodoRequest struct {
	Text string `json:"text"`
}

func (r *CreateTodoRequest) Validate() error {
	if r.Text == "" {
		return errParamIsRequired("text", "string")
	}
	return nil
}

// CreateTodo
type UpdateTodoRequest struct {
	Text string `json:"text"`
}

func (r *UpdateTodoRequest) Validate() error {
	if r.Text == "" {
		return errParamIsRequired("text", "string")
	}
	return nil
}
