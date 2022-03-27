// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Body   string `json:"body"`
	Email  string `json:"email"`
}

type Post struct {
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	User     *User  `json:"user"`
	Comments int    `json:"comments"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
