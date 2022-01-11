package models

type User struct {
	ID			uint	`json:"id" gorm:"primary"`
	Name 		string	`json:"name"`
	Email 		string	`json:"email" gorm:"unique"`
	Password 	string	`json:"password"`
}

type WishItem struct {
	Id				uint	`json:"id"`
	Title			string	`json:"title"`
	Price			float32	`json:"price"`
	Link			string	`json:"link"`
	IsPurchased		bool	`json:"isPurchased"`
	UserID			uint	`json:"-"`
	User			User	`json:"-"`
}