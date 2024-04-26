package model

type Comment struct {
	Code      string `json:"code"`
	BlogCode  string `json:"blog_code"`
	Content   string `json:"content"`
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
}
