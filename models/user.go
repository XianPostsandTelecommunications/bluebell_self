package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}

type Captcha struct {
	Id           string `json:"id"`
	Base64Blob   string `json:"base_64_blob"`
	VertifyValue string `json:"vertify_value"`
}
