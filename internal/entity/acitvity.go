package entity

type Activity struct {
	ActivityID int64  `json:"id"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
}

type CreateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}
