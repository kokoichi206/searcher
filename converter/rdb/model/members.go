package model

type Member struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	EnglishName   string `json:"english_name"`
	Kana          string `json:"kana"`
	Cate          string `json:"cate"`
	Img           string `json:"img"`
	Link          string `json:"link"`
	Pick          string `json:"pick"`
	God           string `json:"god"`
	Under         string `json:"under"`
	Birthday      string `json:"birthday"`
	Blood         string `json:"blood"`
	Constellation string `json:"constellation"`
	Graduation    string `json:"graduation"`
	Groupcode     string `json:"groupcode"`
}
