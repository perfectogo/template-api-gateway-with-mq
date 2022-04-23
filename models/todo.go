package models

import "time"

type Todo struct {
	TodoId    string    `json:"todoId"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type RespTodo struct {
	TodoId    string    `json:"todoId"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CrUpTodo struct {
	TodoId   string `json:"todoId"`
	Title    string `json:"title"`
	Notes    string `json:"notes"`
	Priority string `json:"priority"`
}

type RespTodoList struct {
	Todos RespTodo `json:"todos"`
	Count int64    `json:"count"`
}

type ReqById struct {
	TodoId string `json:"todoId"`
}

type ReqList struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}
