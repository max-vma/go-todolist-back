package dto

type CreateTodoRequest struct {
	Text string `json:"text" binding:"required"`
}

type UpdateTodoRequest struct {
	Text string `json:"text" binding:"required"`
}
