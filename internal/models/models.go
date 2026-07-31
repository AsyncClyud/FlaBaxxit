package models

type User struct {
	Id              int     `json:"Id"`
	Username        string  `json:"Username"`
	Password        string  `json:"Password"`
	Bio             *string `json:"Bio"`
	Created_at      *string `json:"Created_at"`
	Profile_pic int `json:"Profile_pic"`
	Turnstile_token string  `json:"Turnstile_token"`
}

type Article struct {
	Id              int     `json:"Id"`
	Title           string  `json:"Title"`
	Content         string  `json:"Content"`
	Created_at      *string `json:"Created_at"`
	Author          int     `json:"Author"`
	Turnstile_token string  `json:"Turnstile_token"`
}

type ArticleWithAuthor struct {
	Id              int     `json:"Id"`
	Title           string  `json:"Title"`
	Content         string  `json:"Content"`
	Created_at      *string `json:"Created_at"`
	Author_Id       int     `json:"Author_Id"`
	Author_Username string  `json:"Author_Username"`
	Author_Avatar   int     `json:"Author_Avatar"`
}

type Comment struct {
	Id              int     `json:"Id"`
	Comment_content string  `json:"Comment_content"`
	Created_at      *string `json:"Created_at"`
	Post_id         int     `json:"Post_id"`
	Author_Id          int     `json:"Author_Id"`
	Author_Username string  `json:"Author_Username"`
	Author_Avatar   int     `json:"Author_Avatar"`
	Turnstile_token string  `json:"Turnstile_token"`
}

type Message struct {
	Message string `json:"Message"`
}

type NewPassword struct {
	Old_Password string `json:"Old_Password"`
	New_Password string `json:"New_Password"`
}
