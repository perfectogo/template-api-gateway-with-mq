package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	uid "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/perfectogo/template-api-gateway-with-mq/models"
	"github.com/streadway/amqp"
)

//
func (h *Handler) Ping(ctx *gin.Context) {
	h.handleSuccessResponse(ctx, 200, "ok", "pong")
}

//
func (h *Handler) PostTodo(ctx *gin.Context) {
	var (
		entity models.CrUpTodo
	)

	err := ctx.ShouldBindJSON(&entity)
	if err != nil {
		h.handleErrorResponse(ctx, 400, "parse error", err)
		return
	}
	id, err := uid.NewV4()
	if err != nil {
		h.handleErrorResponse(ctx, 400, "uuid generating error", err)
		return
	}
	entity.TodoId = id.String()
	//
	b, err := json.Marshal(entity)
	if err != nil {
		h.handleErrorResponse(ctx, 500, "marshalling error", err)
		return
	}

	uuid, err := uuid.NewRandom()

	if err != nil {
		h.handleErrorResponse(ctx, 500, "server error", err)
		return
	}

	RMQHeaders := make(amqp.Table)

	err = h.rmq.Push("v1.todo", "v1.todo.create", amqp.Publishing{
		Headers:      RMQHeaders,
		Type:         "JSON",
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		ReplyTo:      "v1.response",
		MessageId:    uuid.String(),
		Body:         b,
	})

	if err != nil {
		h.handleErrorResponse(ctx, 500, "server error", err)
		return
	}

	h.handleSuccessResponse(ctx, 201, "todo is being created", entity)

}

//
func (h *Handler) GetTodoByID(ctx *gin.Context) {
	err := h.makeProxy(ctx, h.cfg.TodoServiceURL, ctx.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(ctx, 500, "proxy error", err)
	}
}
