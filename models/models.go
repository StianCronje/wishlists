package models

type User struct {
	ID			uint	`json:"id" gorm:"primary"`
	Name 		string	`json:"name"`
	Email 		string	`json:"email" gorm:"unique"`
	Password 	[]byte	`json:"-"`
}

type WishItem struct {
	Id				string
	Title			string
	Price			float32
	Link			string
	IsPurchased		bool
	UserID			uint
	User			User
}