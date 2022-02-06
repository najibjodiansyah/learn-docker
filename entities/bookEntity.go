package entities

type Book struct {
	Id      	int     `json:"id" form:"id"`
	Title 	 	string	`json:"title" form:"title"`
	Description string	`json:"description" form:"description"`
	Publisher 	string	`json:"publisher" form:"publisher"`
}
