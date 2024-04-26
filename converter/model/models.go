package model

type Member struct {
	Birthday   string `json:"birthday"`
	BlogUrl    string `json:"blog_url"`
	BloodType  string `json:"blood_type"`
	Generation string `json:"generation"`
	Height     string `json:"height"`
	ImgUrl     string `json:"img_url"`
	Name       string `json:"name"`
}

type Blog struct {
	ArtiCode  string `json:"arti_code"`
	Cate      string `json:"cate"`
	Code      string `json:"code"`
	Date      string `json:"date"`
	EndTime   string `json:"end_time"`
	Images    string `json:"images"`
	Link      string `json:"link"`
	Member    string `json:"member"`
	StartTime string `json:"start_time"`
	Text      string `json:"text"`
	Title     string `json:"title"`
}

type Comment struct {
	Code       string `json:"code"`
	BlogCode   string `json:"blog_code"`
	BlogTitle  string `json:"blog_title"`
	MemberName string `json:"member_name"`
	Text       string `json:"text"`
	Name       string `json:"name"`
	Date       string `json:"date"`
}
