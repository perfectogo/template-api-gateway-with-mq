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
type TodoQueryParamModel struct {
	Search      string `json:"search"`
	BranchID    string `json:"branch_id"`
	Order       string `json:"order" enums:"id,name,region_id,address,location,vehicle_number,vehicle_model,branch_id,phone,created_at, updated_at"`
	Arrangement string `json:"arrangement" enums:"asc,desc"`
	Offset      int    `json:"offset" default:"0"`
	Limit       int    `json:"limit" default:"10"`
}
