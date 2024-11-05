package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/adpter/input/controller"
)

func InitRoutes(r *gin.RouterGroup, controller controller.WebhookControllerInterface) {
	r.POST("/:id/:token", controller.CreatePR)
}
