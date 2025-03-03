package api

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Users = []User{
	{ID: 1, Name: "Otto"},
	{ID: 2, Name: "Jeaneth"},
}