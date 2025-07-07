package request

type CommentCreateRequest struct {
	Content string `json:"content" validate:"required"`
	PostID  uint   `json:"post_id" validate:"required"`
}
