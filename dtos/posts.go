package dtos

type Post struct {
	Id          int64  `json:"id"`
	PostContent string `json:"post_content"`
	PostImage   string `json:"post_image"`
	PostedBy    string `json:"posted_by"`
}

type PostRequest struct {
	PostContent string `json:"post_content"`
	PostImage   string `json:"post_image"`
	PostedBy    string `json:"posted_by"`
}

type PostResponse struct {
	Id          int64  `json:"id"`
	PostContent string `json:"post_content"`
	PostImage   string `json:"post_image"`
	PostedBy    string `json:"posted_by"`
}
