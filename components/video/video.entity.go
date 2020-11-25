package video

type Video struct {
	Id          string `json:"-"`
	Title       string `json:"title" binding:"min=2,max=20"`
	Description string `json:"description" binding:"required"`
	URL         string `json:"url" binding:"required,url"`
	// UserId      string `json:"userId"`
	// Author      user.User `json:"author"`
}
