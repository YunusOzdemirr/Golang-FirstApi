package models

type User struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
	Profile   string
	Interests []Interest
}
