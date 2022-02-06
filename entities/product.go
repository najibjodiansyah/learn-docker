package entities

type Product struct {
	Id		int     `json:"id" form:"id"`
	UserId 	User	
	Nama 	string	`json:"nama" form:"nama"`
	Harga 	string	`json:"harga" form:"harga"`
}