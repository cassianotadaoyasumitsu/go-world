package main

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `email:"email"`
}

type Hotel struct {
	ID          int    `json:"id"`
	Country     string `json:"country"`
	Destination string `json:"destination"`
	Name        string `json:"hotel"`
}
