package controller

type ArticleCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	AuthorId    int    `json:"author_id"`
}
