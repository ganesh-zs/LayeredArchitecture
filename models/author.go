package models

type Author struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PenName     string `json:"penName"`
	DateOfBirth string `json:"dateOfBirth"`
	Genre       string `json:"genre"`
}
