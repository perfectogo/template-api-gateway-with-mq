package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/perfectogo/template-api-gateway-with-mq/api/handlers/v1"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	r.Use(MaxAllowed(50))
	{
		r.POST("/todo", h.PostTodo)
		r.GET("/todos", h.GetTodoList)
		r.GET("/todos/:id", h.GetTodoByID)
		r.PUT("/todos/:id", h.PutTodo)
		r.DELETE("/todos/:id", h.DeleteTodo)
	}
}

//
func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
		fmt.Println("countRequest: ", countReq)
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
		fmt.Println("countResp: ", countReq)
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}
