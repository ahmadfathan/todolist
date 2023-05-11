package entity

type Todo struct {
	TodoID          int64  `json:"id"`
	ActivityGroupID int64  `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

type CreateTodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

type UpdateTodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive        *bool  `json:"is_active"`
	Priority        string `json:"priority"`
}
