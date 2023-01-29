package models

type Book struct {
	Id                int     `json:"id,omitempty"`
	ISBN              int     `json:"isbn"`
	Title             string  `json:"title"`
	Genre             string  `json:"genre"`
	Publication       string  `json:"publication"`
	YearOfPublication int     `json:"yearOfPublication"`
	AuthorId          int     `json:"authorId"`
	BookAuthor        *Author `json:"bookAuthor,omitempty"`
}
