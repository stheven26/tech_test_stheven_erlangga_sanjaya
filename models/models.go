package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type Data struct {
	ID             string `json:"_id"`
	Name           string `json:"name"`
	Balance        int    `json:"balance"`
	Transportation string `json:"transportation"`
}
