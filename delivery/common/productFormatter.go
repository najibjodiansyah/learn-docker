package common

type ProductRequestFormat struct {
	Nama 	string	`json:"nama" form:"nama"`
	Harga 	string	`json:"harga" form:"harga"`
}

type ProductResponseFormat struct {
	Id 			int		`json:"id" form:"id"`
	UserId		int		`json:"userid" form:"userid"`
	Username	string	`json:"username" form:"username"`
	Nama 		string	`json:"nama" form:"nama"`
	Harga 		string	`json:"harga" form:"harga"`
}