package model

type Blog struct {
	Code                   string `json:"code"`
	Title                  string `json:"title"`
	Content                string `json:"content"`
	Timestamp              string `json:"timestamp"`
	StartTime              string `json:"start_time"`
	EndTime                string `json:"end_time"`
	Cate                   string `json:"cate"`
	Link                   string `json:"link"`
	MemberCode             string `json:"member_code"`
	LatestCommentTimestamp string `json:"latest_comment_timestamp"`
}
